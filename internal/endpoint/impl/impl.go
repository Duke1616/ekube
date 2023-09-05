package impl

import (
	v1 "ekube/api/pb/endpoint/v1"
	"ekube/config"
	"ekube/internal/endpoint"
	"ekube/internal/endpoint/data"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

var _ endpoint.Service = &service{}

type service struct {
	log      logger.Logger
	endpoint endpoint.Service

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

	s.endpoint = ioc.GetInternalApp(endpoint.AppName).(endpoint.Service)
	return nil
}

func (s *service) Name() string {
	return endpoint.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
