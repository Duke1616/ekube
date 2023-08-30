package controller

import (
	"context"
	"ekube/cmd/controller/option"
	"ekube/cmd/signals"
	"ekube/config"
	"ekube/pkg/apis"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func NewControllerCommand(s *option.KubeSphereControllerManagerOptions) *cobra.Command {
	Cmd := &cobra.Command{
		Use:   "controller",
		Short: "ekube controller服务",
		Long:  "ekube controller服务",
		RunE: func(cmd *cobra.Command, args []string) error {
			//if errs := s.Validate(); len(errs) != 0 {
			//	return nil
			//}

			return Run(s, config.WatchConfigChange(), signals.SetupSignalHandler())
		},
		SilenceUsage: true,
	}

	return Cmd
}

func Run(s *option.KubeSphereControllerManagerOptions, configCh <-chan config.Config, ctx context.Context) error {
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
			s.MergeConfig(&cfg)
			ictx, cancelFunc = context.WithCancel(context.TODO())
			go func() {
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

func run(s *option.KubeSphereControllerManagerOptions, ctx context.Context) error {

	kubernetesClient, err := client.NewKubernetesClient(s.KubernetesOption)
	if err != nil {
		klog.Errorf("Failed to create kubernetes clientset %v", err)
		return err
	}

	informerFactory := informer.NewInformerFactories(
		kubernetesClient.Kubernetes(),
	)

	mgrOptions := manager.Options{
		CertDir: s.WebhookCertDir,
		Port:    8443,
	}

	if s.LeaderElect {
		mgrOptions = manager.Options{
			CertDir:                 s.WebhookCertDir,
			Port:                    8443,
			LeaderElection:          s.LeaderElect,
			LeaderElectionNamespace: "ekube-system",
			LeaderElectionID:        "ks-controller-manager-leader-election",
			LeaseDuration:           &s.LeaderElection.LeaseDuration,
			RetryPeriod:             &s.LeaderElection.RetryPeriod,
			RenewDeadline:           &s.LeaderElection.RenewDeadline,
		}
	}

	klog.V(0).Info("setting up manager")
	ctrl.SetLogger(klog.NewKlogr())
	// Use 8443 instead of 443 cause we need root permission to bind port 443
	mgr, err := manager.New(kubernetesClient.Config(), mgrOptions)
	if err != nil {
		klog.Fatalf("unable to set up overall controller manager: %v", err)
	}

	if err = apis.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatalf("unable add APIs to scheme: %v", err)
	}

	// register common meta types into schemas.
	metav1.AddToGroupVersion(mgr.GetScheme(), metav1.SchemeGroupVersion)

	// install all controllers
	if err = addAllControllers(mgr,
		kubernetesClient,
		informerFactory,
		s,
		ctx.Done()); err != nil {
		klog.Fatalf("unable to register controllers to the manager: %v", err)
	}

	// Start cache data after all informer is registered
	klog.V(0).Info("Starting cache resource from apiserver...")
	informerFactory.Start(ctx.Done())

	// Starting the controllers
	klog.V(0).Info("Starting the controllers.")
	if err = mgr.Start(ctx); err != nil {
		klog.Fatalf("unable to run the manager: %v", err)
	}

	return nil
}
