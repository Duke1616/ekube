package start

import (
	"context"
	"ekube/cmd/start/option"
	"ekube/cmd/start/signals"
	"ekube/conf"
	// 注册所有服务
	_ "ekube/internal/all"
	"github.com/spf13/cobra"
)

func NewAPIServerCommand(s *option.ServerRunOptions) *cobra.Command {
	Cmd := &cobra.Command{
		Use:   "start",
		Short: "ekube API服务",
		Long:  "ekube API服务",
		RunE: func(cmd *cobra.Command, args []string) error {
			if errs := s.Validate(); len(errs) != 0 {
				return nil
			}

			return Run(s, conf.WatchConfigChange(), signals.SetupSignalHandler())
		},
		SilenceUsage: true,
	}

	return Cmd
}

func Run(s *option.ServerRunOptions, configCh <-chan conf.Config, ctx context.Context) error {
	ictx, cancelFunc := context.WithCancel(context.TODO())
	errCh := make(chan error)
	defer close(errCh)
	go func() {
		if err := run(s, ictx); err != nil {
			errCh <- err
		}
	}()

	// ctx 控制整个程序的生命周期
	// ictx 通过viper配置文件变更，重启服务
	for {
		select {
		case <-ctx.Done():
			cancelFunc()
			return nil
		case cfg := <-configCh:
			cancelFunc()
			s.Config = &cfg

			ictx, cancelFunc = context.WithCancel(context.TODO())
			go func() {
				if errs := s.Validate(); len(errs) != 0 {
					for _, err := range errs {
						errCh <- err
					}
				}
				if err := run(s, ictx); err != nil {
					errCh <- err
				}
			}()
		case err := <-errCh:
			cancelFunc()
			return err
		}
	}
}

func run(s *option.ServerRunOptions, ctx context.Context) error {
	api, err := s.NewAPIServer()
	if err != nil {
		return err
	}

	api.PrepareRun(ctx.Done())

	api.Start(ctx)

	return nil
}

//// Cmd represents the start command
//var Cmd = &cobra.Command{
//	Use:   "start",
//	Short: "ekube API服务",
//	Long:  "ekube API服务",
//	Run: func(cmd *cobra.Command, args []string) {
//		conf := conf.C()
//		// 启动服务
//		ch := make(chan os.Signal, 1)
//		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
//
//		// 初始化服务
//		svr, err := newService(conf)
//		cobra.CheckErr(err)
//
//		// 启动服务
//		svr.start()
//	},
//}

//func newService(conf *conf.Config) (*service, error) {
//	http := protocol.NewHTTPService()
//
//	client, err := conf.Etcd.KubernetesClient()
//	if err != nil {
//		panic(err)
//	}
//
//	r, err := etcd.NewRegistry(client)
//	if err != nil {
//		panic(err)
//	}
//
//	grpc, err := protocol.NewGRPCService("ekube", protocol.ServerWithRegistry(r))
//	if err != nil {
//		panic(err)
//	}
//
//	// 处理信号量
//	ch := make(chan os.Signal, 1)
//	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
//	svr := &service{
//		http: http,
//		grpc: grpc,
//		log:  zap.L().Named("cli"),
//		ch:   ch,
//	}
//
//	return svr, nil
//}

//type service struct {
//	http *protocol.HTTPService
//	grpc *protocol.GRPCService
//	ch   chan os.Signal
//	log  logger.Logger
//}
//
//func (s *service) start(ctx context.Context) {
//	s.log.Infof("loaded grpc app: %s", ioc.LoadedGrpcApp())
//	s.log.Infof("loaded internal app: %s", ioc.LoadedInternalApp())
//	s.log.Infof("loaded restful app: %s", ioc.LoadedRestfulApp())
//
//	go s.grpc.Start()
//	go s.http.Start(ctx)
//	s.waitSign(s.ch)
//}
//
//func (s *service) waitSign(sign chan os.Signal) {
//	for sg := range sign {
//		switch v := sg.(type) {
//		default:
//			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
//
//			if err := s.grpc.Close(); err != nil {
//				s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
//			} else {
//				s.log.Info("grpc service stop complete")
//			}
//
//			if err := s.http.Stop(); err != nil {
//				s.log.Errorf("http graceful shutdown err: %s, force exit", err)
//			} else {
//				s.log.Infof("http service stop complete")
//			}
//
//			// 关闭依赖的全景配置对象
//			conf.C().Shutdown(context.Background())
//			return
//		}
//	}
//}
