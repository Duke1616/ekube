package protocol

import (
	"context"
	"ekube/conf"
	"ekube/protocol/ioc"
	"ekube/swagger"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	//r := restful.DefaultContainer
	r := restful.NewContainer()
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)

	// CORS中间件
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}
	r.Filter(cors.Filter)
	//// trace中间件
	//filter := otelrestful.OTelFilter(version.ServiceName)
	//restful.DefaultContainer.Filter(filter)
	//// 认证中间件
	//r.Filter(middleware.RestfulServerInterceptor())

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTP.Addr(),
		Handler:           r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("server.http"),
		c:      conf.C(),
		//endpoint:   c.Endpoint(),
		apiDocPath: "/apidocs.json",
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *restful.Container
	l      logger.Logger
	c      *conf.Config
	server *http.Server

	//endpoint   endpoint.RPCClient
	apiDocPath string
}

func (s *HTTPService) PathPrefix() string {
	return s.c.App.HTTPPrefix()
}

// Start 启动服务
func (s *HTTPService) Start(ctx context.Context) {
	go func() {
		<-ctx.Done()
		_ = s.Stop()
	}()

	// 装置子服务路由
	ioc.LoadRestfulApp(s.PathPrefix(), s.r)

	config := restfulspec.Config{
		WebServices:                   s.r.RegisteredWebServices(), // you control what services are visible
		APIPath:                       s.apiDocPath,
		PostBuildSwaggerObjectHandler: swagger.Docs,
		DefinitionNameHandler: func(name string) string {
			if name == "state" || name == "sizeCache" || name == "unknownFields" {
				return ""
			}
			return name
		},
	}

	s.r.Add(restfulspec.NewOpenAPIService(config))

	s.l.Infof("Get the API Doc using http://%s%s", s.c.App.HTTP.Addr(), s.apiDocPath)

	// 注册路由条目
	s.RegistryEndpoint()

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
			return
		}
		s.l.Errorf("start service error, %s", err.Error())
	}
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}

func (s *HTTPService) RegistryEndpoint() {
	// 注册服务权限条目
	s.l.Info("start registry endpoints ...")

	//entries := []*endpoint.Entry{}
	//wss := s.r.RegisteredWebServices()
	//for i := range wss {
	//	es := endpoint.TransferRoutesToEntry(wss[i].Routes())
	//	entries = append(entries, endpoint.GetPRBACEntry(es)...)
	//}
	//
	//req := endpoint.NewRegistryRequest(version.Short(), entries)
	//_, err := s.endpoint.RegistryEndpoint(context.Background(), req)
	//if err != nil {
	//	s.l.Warnf("registry endpoints error, %s", err)
	//} else {
	//	s.l.Debug("service endpoints registry success")
	//}
}
