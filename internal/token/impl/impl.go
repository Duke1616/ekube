package impl

import (
	v1 "ekube/api/pb/token/v1"
	"ekube/config"
	"ekube/internal/token"
	"ekube/internal/token/data"
	"ekube/internal/token/provider"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	_ "ekube/internal/token/provider/all"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

var _ token.Service = &service{}

type service struct {
	log   logger.Logger
	token token.Service

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

	s.token = ioc.GetInternalApp(token.AppName).(token.Service)

	// 初始化所有的auth provider
	if err = provider.Init(); err != nil {
		return err
	}
	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
