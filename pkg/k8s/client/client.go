package client

import (
	"ekube/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client interface {
	Kubernetes() kubernetes.Interface
	Config() *rest.Config
}

type kubernetesClient struct {
	k8s kubernetes.Interface

	config *rest.Config
}

func NewKubernetesClient(option *config.Kubernetes) (Client, error) {
	conf, err := clientcmd.BuildConfigFromFlags("", option.KubeConfig)
	if err != nil {
		return nil, err
	}

	conf.QPS = option.QPS
	conf.Burst = option.Burst

	if err != nil {
		return nil, err
	}

	var k kubernetesClient
	k.k8s, err = kubernetes.NewForConfig(conf)

	if err != nil {
		return nil, err
	}

	k.config = conf

	return &k, nil
}

func (k *kubernetesClient) Kubernetes() kubernetes.Interface {
	return k.k8s
}

func (k *kubernetesClient) Config() *rest.Config {
	return k.config
}
