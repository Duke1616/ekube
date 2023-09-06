package impl

import (
	v1 "ekube/api/pb/role/v1"
	"ekube/config"
	"ekube/internal/role"
	"ekube/internal/role/data"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

var _ role.Service = &service{}

type service struct {
	log  logger.Logger
	role role.Service

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
	s.role = ioc.GetInternalApp(role.AppName).(role.Service)
	return nil
}

func (s *service) Name() string {
	return role.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
