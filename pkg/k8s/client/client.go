package client

import (
	"ekube/conf"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client interface {
	Kubernetes() kubernetes.Interface
}

type kubernetesClient struct {
	k8s kubernetes.Interface
}

func NewKubernetesClient(option *conf.Kubernetes) (Client, error) {
	config, err := clientcmd.BuildConfigFromFlags("", option.KubeConfig)
	if err != nil {
		return nil, err
	}

	config.QPS = option.QPS
	config.Burst = option.Burst

	if err != nil {
		return nil, err
	}

	var k kubernetesClient
	k.k8s, err = kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return &k, nil
}

func (k *kubernetesClient) Kubernetes() kubernetes.Interface {
	return k.k8s
}
