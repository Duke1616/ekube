package deployment

import (
	"ekube/pkg/apiserver/query"
	"ekube/pkg/k8s/resources"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
)

type deploymentsGetter struct {
	sharedInformers informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) resources.Interface {
	return &deploymentsGetter{sharedInformers: sharedInformers}
}

func (d *deploymentsGetter) Get(namespace, name string) (runtime.Object, error) {
	return d.sharedInformers.Apps().V1().Deployments().Lister().Deployments(namespace).Get(name)
}

func (d *deploymentsGetter) List(namespace string, query *query.Query) (*resources.ListResult, error) {
	deployments, err := d.sharedInformers.Apps().V1().Deployments().Lister().Deployments(namespace).List(query.Selector())
	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, deployment := range deployments {
		result = append(result, deployment)
	}

	return nil, nil
	//return v1alpha3.DefaultList(result, query, d.compare, d.filter), nil
}
