package protocol

import (
	"context"
	"ekube/conf"
	"ekube/protocol/ioc"
	"ekube/registry"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type ServerOption func(server *GRPCService)

// GRPCService grpc服务
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

// Start 启动GRPC服务
func (s *GRPCService) Start(ctx context.Context) {
	go func() {
		<-ctx.Done()
		_ = s.Close()
	}()

	ioc.LoadGrpcApp(s.Server)

	listener, err := net.Listen("tcp", conf.C().App.GRPC.Addr())
	if err != nil {
		return
	}

	s.listener = listener

	if s.registry != nil {
		ctx, cancel := context.WithTimeout(context.Background(), s.registryTimeout)

		defer cancel()
		err = s.registry.Registry(ctx, registry.ServiceInstance{
			Name:    s.Name,
			Address: listener.Addr().String(),
		})
		if err != nil {
			return
		}
	}

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPC.Addr())
	if err = s.Serve(listener); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

func (s *GRPCService) Close() error {
	if s.registry != nil {
		err := s.registry.Close()
		if err != nil {
			return err
		}
	}

	s.GracefulStop()
	return nil
}
