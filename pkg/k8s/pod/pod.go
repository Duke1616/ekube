package pod

import (
	"ekube/pkg/apiserver/query"
	"ekube/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
)

type podsGetter struct {
	informer informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) k8s.Interface {
	return &podsGetter{informer: sharedInformers}
}

func (p *podsGetter) Get(namespace, name string) (runtime.Object, error) {
	return p.informer.Core().V1().Pods().Lister().Pods(namespace).Get(name)
}

func (p *podsGetter) List(namespace string, query *query.Query) (*k8s.ListResult, error) {
	pods, err := p.informer.Core().V1().Pods().Lister().Pods(namespace).List(query.Selector())

	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, pod := range pods {
		result = append(result, pod)
	}

	return k8s.DefaultList(result, nil), nil
}
