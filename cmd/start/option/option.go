package option

import (
	"ekube/conf"
	"ekube/pkg/apiserver"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/client"
	"k8s.io/client-go/util/homedir"
	"os"
	"os/user"
	"path"
	"sync"
)

func (s *ServerRunOptions) NewAPIServer() (*apiserver.APIServer, error) {
	apiServer := &apiserver.APIServer{}
	//apiServer := newService(conf.C())

	kClient, err := client.NewKubernetesClient(s.KubernetesOption)
	if err != nil {
		return nil, err
	}

	apiServer.KubernetesClient = kClient

	informerFactory := informer.NewInformerFactories(kClient.Kubernetes())

	apiServer.InformerFactory = informerFactory
	apiServer.Config = s.KubernetesOption

	return apiServer, nil
}

type ServerRunOptions struct {
	ConfigFile       string
	schemeOnce       sync.Once
	KubernetesOption *conf.Kubernetes
	*conf.Config
	DebugMode bool
}

func NewServerRunOptions() *ServerRunOptions {
	s := &ServerRunOptions{
		KubernetesOption: NewKubernetesOptions(),
		schemeOnce:       sync.Once{},
	}

	return s
}

func NewKubernetesOptions() (option *conf.Kubernetes) {
	option = &conf.Kubernetes{
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
