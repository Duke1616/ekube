package informer

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"reflect"
	"time"
)

const defaultResync = 600 * time.Second

// InformerFactory callers should check if the return value is nil
type InformerFactory interface {
	KubernetesSharedInformerFactory() informers.SharedInformerFactory

	// Start shared informer factory one by one if they are not nil
	Start(stopCh <-chan struct{})
}

type GenericInformerFactory interface {
	Start(stopCh <-chan struct{})
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

type informerFactories struct {
	informerFactory informers.SharedInformerFactory
}

func NewInformerFactories(client kubernetes.Interface) InformerFactory {
	factory := &informerFactories{}
	if client != nil {
		factory.informerFactory = informers.NewSharedInformerFactory(client, defaultResync)
	}

	return factory
}

func (f *informerFactories) Start(stopCh <-chan struct{}) {
	if f.informerFactory != nil {
		f.informerFactory.Start(stopCh)
	}
}

func (f *informerFactories) KubernetesSharedInformerFactory() informers.SharedInformerFactory {
	return f.informerFactory
}
