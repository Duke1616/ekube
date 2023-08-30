package apiserver

import (
	"context"
	"ekube/config"
	proxyAPI "ekube/internal/proxy/api"
	terminalAPI "ekube/internal/terminal/api"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/client"
	"ekube/protocol"
	"ekube/protocol/ioc"
	"ekube/registry/etcd"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
)

type APIServer struct {
	ServerCount int

	Server *http.Server

	Config *config.Kubernetes

	http *protocol.HTTPService

	grpc *protocol.GRPCService

	KubernetesClient client.Client

	InformerFactory informer.InformerFactory

	log logger.Logger
}

func (s *APIServer) newService() *APIServer {
	httpServer := protocol.NewHTTPService()

	etcdClient, err := config.C().Etcd.Client()
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewRegistry(etcdClient)
	if err != nil {
		panic(err)
	}

	grpc, err := protocol.NewGRPCService("ekube", protocol.ServerWithRegistry(r))
	if err != nil {
		panic(err)
	}

	s.http = httpServer
	s.grpc = grpc
	s.log = zap.L().Named("cli")

	return s
}

func (s *APIServer) PrepareRun(stopCh <-chan struct{}) {
	s.newService()

	urlruntime.Must(proxyAPI.Handler.AddToContainer(s.InformerFactory, s.KubernetesClient.Kubernetes()))
	urlruntime.Must(terminalAPI.Handler.AddToContainer(s.KubernetesClient.Kubernetes(), s.KubernetesClient.Config(), config.C().TerminalOption))
	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
	}

	for _, v := range gvrs {
		// 创建 informer
		_, err := s.InformerFactory.KubernetesSharedInformerFactory().ForResource(v)
		if err != nil {
			return
		}
	}

	s.InformerFactory.Start(stopCh)
	s.InformerFactory.KubernetesSharedInformerFactory().WaitForCacheSync(stopCh)
}

func (s *APIServer) Start(ctx context.Context) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	s.log.Infof("loaded grpc app: %s", ioc.LoadedGrpcApp())
	s.log.Infof("loaded internal app: %s", ioc.LoadedInternalApp())
	s.log.Infof("loaded restful app: %s", ioc.LoadedRestfulApp())

	go s.grpc.Start(ctx)
	go s.http.Start(ctx)
	s.waitSign(ch)
}

func (s *APIServer) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())

			if err := s.grpc.Close(); err != nil {
				s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Info("grpc service stop complete")
			}

			if err := s.http.Stop(); err != nil {
				s.log.Errorf("http graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Infof("http service stop complete")
			}

			// 关闭依赖的全景配置对象
			config.C().Shutdown(context.Background())
		}
	}
}
