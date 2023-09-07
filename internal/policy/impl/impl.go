package impl

import (
	v1 "ekube/api/pb/policy/v1"
	"ekube/config"
	"ekube/internal/endpoint"
	"ekube/internal/namespace"
	"ekube/internal/policy"
	"ekube/internal/policy/data"
	"ekube/internal/role"
	"ekube/internal/workspace"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

var _ policy.Service = &service{}

type service struct {
	log    logger.Logger
	policy policy.Service

	data      *data.Data
	workspace workspace.Service
	namespace namespace.Service
	endpoint  endpoint.Service
	role      role.Service
	v1.UnimplementedRPCServer

	encryptKey string
}

func (s *service) Config() error {
	db, err := config.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.encryptKey = config.C().App.EncryptKey
	s.data = data.NewData(db, s.Name())
	s.log = zap.L().Named(s.Name())

	s.policy = ioc.GetInternalApp(policy.AppName).(policy.Service)
	s.workspace = ioc.GetInternalApp(workspace.AppName).(workspace.Service)
	s.endpoint = ioc.GetInternalApp(endpoint.AppName).(endpoint.Service)
	s.namespace = ioc.GetInternalApp(namespace.AppName).(namespace.Service)
	s.role = ioc.GetInternalApp(role.AppName).(role.Service)
	return nil
}

func (s *service) Name() string {
	return policy.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
