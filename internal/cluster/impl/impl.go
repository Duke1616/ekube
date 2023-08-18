package impl

import (
	v1 "ekube/api/pb/cluster/v1"
	"ekube/conf"
	"ekube/internal/cluster"
	"ekube/internal/cluster/data"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryGrpcApp(&service{})
	ioc.RegistryInternalApp(&service{})
}

type service struct {
	log     logger.Logger
	cluster cluster.Service

	data *data.Data
	v1.UnimplementedRPCServer

	encryptKey string
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.encryptKey = conf.C().App.EncryptKey

	s.data = data.NewData(db, s.Name())

	s.log = zap.L().Named(s.Name())
	s.cluster = ioc.GetInternalApp(cluster.AppName).(cluster.Service)
	return nil
}

func (s *service) Name() string {
	return cluster.AppName
}

func (s *service) Registry(server *grpc.Server) {
	v1.RegisterRPCServer(server, s)
}
