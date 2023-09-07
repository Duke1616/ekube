package impl

import (
	v1 "ekube/api/pb/workspace/v1"
	"ekube/config"
	"ekube/internal/cluster"
	"ekube/internal/workspace"
	"ekube/internal/workspace/data"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

var _ workspace.Service = &service{}

type service struct {
	log       logger.Logger
	workspace workspace.Service
	cluster   cluster.Service

	data *data.Data
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

	s.workspace = ioc.GetInternalApp(workspace.AppName).(workspace.Service)
	s.cluster = ioc.GetInternalApp(cluster.AppName).(cluster.Service)
	return nil
}

func (s *service) Name() string {
	return workspace.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
