package ioc

import (
	"ekube/conf"
	"ekube/registry"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type ServerOption func(server *GRPCService)

type GRPCService struct {
	Name string
	*grpc.Server
	l logger.Logger
	c *conf.Config

	registry        registry.Registry
	registryTimeout time.Duration

	listener net.Listener
}

// NewGRPCService todo
func NewGRPCService(name string, opts ...ServerOption) (*GRPCService, error) {
	res := &GRPCService{
		Name:            name,
		Server:          grpc.NewServer(),
		l:               zap.L().Named("server.grpc"),
		c:               conf.C(),
		registryTimeout: time.Second * 10,
	}

	for _, opt := range opts {
		opt(res)
	}

	return res, nil
}

func ServerWithRegistry(r registry.Registry) ServerOption {
	return func(server *GRPCService) {
		server.registry = r
	}
}
