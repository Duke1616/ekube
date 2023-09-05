package deployment

import (
	"ekube/pkg/apiserver/query"
	"ekube/pkg/k8s/resources"
	v1 "k8s.io/api/apps/v1"
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
	// 处理过滤、排序逻辑
	options := resources.ListOptions{
		CompareFunc: d.compare,
		FilterFunc:  d.filter,
	}

	return resources.DefaultList(result, query, options), nil
}

func (d *deploymentsGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {
	leftPod, ok := left.(*v1.Deployment)
	if !ok {
		return false
	}

	rightPod, ok := right.(*v1.Deployment)
	if !ok {
		return false
	}

	return resources.DefaultObjectMetaCompare(leftPod.ObjectMeta, rightPod.ObjectMeta, field)
}

func (d *deploymentsGetter) filter(object runtime.Object, filter query.Filter) bool {
	_, ok := object.(*v1.Deployment)

	if !ok {
		return false
	}

	return true
}
