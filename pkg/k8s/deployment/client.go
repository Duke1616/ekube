package deployment

import (
	"k8s.io/client-go/kubernetes"

	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func NewDeployment(cs kubernetes.Interface) *Client {
	return &Client{
		appsv1: cs.AppsV1(),
	}
}

type Client struct {
	appsv1 appsv1.AppsV1Interface
}
