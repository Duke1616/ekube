package option

import (
	"ekube/config"
	"fmt"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/util/homedir"
	"os"
	"os/user"
	"path"
	"time"
)

type KubeSphereControllerManagerOptions struct {
	KubernetesOption *config.Kubernetes

	LeaderElect    bool
	LeaderElection *leaderelection.LeaderElectionConfig
	WebhookCertDir string

	// KubeSphere is using sigs.k8s.io/application as fundamental object to implement Application Management.
	// There are other projects also built on sigs.k8s.io/application, when KubeSphere installed along side
	// them, conflicts happen. So we leave an option to only reconcile applications  matched with the given
	// selector. Default will reconcile all applications.
	//    For example
	//      "kubesphere.io/creator=" means reconcile applications with this label key
	//      "!kubesphere.io/creator" means exclude applications with this key
	ApplicationSelector string

	// ControllerGates is the list of controller gates to enable or disable controller.
	// '*' means "all enabled by default controllers"
	// 'foo' means "enable 'foo'"
	// '-foo' means "disable 'foo'"
	// first item for a particular name wins.
	//     e.g. '-foo,foo' means "disable foo", 'foo,-foo' means "enable foo"
	// * has the lowest priority.
	//     e.g. *,-foo, means "disable 'foo'"
	ControllerGates []string

	// Enable gops or not.
	GOPSEnabled bool
}

func NewKubeSphereControllerManagerOptions() *KubeSphereControllerManagerOptions {
	s := &KubeSphereControllerManagerOptions{
		LeaderElection: &leaderelection.LeaderElectionConfig{
			LeaseDuration: 30 * time.Second,
			RenewDeadline: 15 * time.Second,
			RetryPeriod:   5 * time.Second,
		},
		LeaderElect:         false,
		WebhookCertDir:      "",
		ApplicationSelector: "",
		ControllerGates:     []string{"*"},
		KubernetesOption:    NewKubernetesOptions(),
	}

	return s
}

func NewKubernetesOptions() (option *config.Kubernetes) {
	option = &config.Kubernetes{
		QPS:   1e6,
		Burst: 1e6,
	}

	// make it be easier for those who wants to run api-server locally
	homePath := homedir.HomeDir()
	if homePath == "" {
		// try os/user.HomeDir when $HOME is unset.
		if u, err := user.Current(); err == nil {
			homePath = u.HomeDir
		}
	}

	userHomeConfig := path.Join(homePath, ".kube/config")
	if _, err := os.Stat(userHomeConfig); err == nil {
		option.KubeConfig = userHomeConfig
	}
	return
}

func (s *KubeSphereControllerManagerOptions) IsControllerEnabled(name string) bool {
	hasStar := false
	for _, ctrl := range s.ControllerGates {
		if ctrl == name {
			return true
		}
		if ctrl == "-"+name {
			return false
		}
		if ctrl == "*" {
			hasStar = true
		}
	}

	return hasStar
}

// MergeConfig merge new config without validation
// When misconfigured, the app should just crash directly
func (s *KubeSphereControllerManagerOptions) MergeConfig(cfg *config.Config) {
	fmt.Print(cfg.Kubernetes)
	s.KubernetesOption = NewKubernetesOptions()
}
