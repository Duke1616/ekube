package pod

import (
	"ekube/pkg/apiserver/query"
	"ekube/pkg/k8s/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
)

type podsGetter struct {
	informer informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) resources.Interface {
	return &podsGetter{informer: sharedInformers}
}

func (p *podsGetter) Get(namespace, name string) (runtime.Object, error) {
	return p.informer.Core().V1().Pods().Lister().Pods(namespace).Get(name)
}

func (p *podsGetter) List(namespace string, query *query.Query) (*resources.ListResult, error) {
	pods, err := p.informer.Core().V1().Pods().Lister().Pods(namespace).List(query.Selector())

	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, pod := range pods {
		result = append(result, pod)
	}

	// 处理过滤、排序逻辑
	options := resources.ListOptions{
		CompareFunc: p.compare,
		FilterFunc:  p.filter,
	}

	return resources.DefaultList(result, query, options), nil
}

func (p *podsGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {
	leftPod, ok := left.(*corev1.Pod)
	if !ok {
		return false
	}

	rightPod, ok := right.(*corev1.Pod)
	if !ok {
		return false
	}

	return resources.DefaultObjectMetaCompare(leftPod.ObjectMeta, rightPod.ObjectMeta, field)
}

func (p *podsGetter) filter(object runtime.Object, filter query.Filter) bool {
	return true
}
