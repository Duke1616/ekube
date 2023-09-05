package resource

import (
	"ekube/pkg/apiserver/query"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/resources"
	"ekube/pkg/k8s/resources/deployment"
	"ekube/pkg/k8s/resources/pod"
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrResourceNotSupported = errors.New("resource is not supported")

type ResourceGetter struct {
	clusterResourceGetters    map[schema.GroupVersionResource]resources.Interface
	namespacedResourceGetters map[schema.GroupVersionResource]resources.Interface
}

func NewResourceGetter(factory informer.InformerFactory) *ResourceGetter {
	namespacedResourceGetters := make(map[schema.GroupVersionResource]resources.Interface)
	clusterResourceGetters := make(map[schema.GroupVersionResource]resources.Interface)

	namespacedResourceGetters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}] = pod.New(factory.KubernetesSharedInformerFactory())
	namespacedResourceGetters[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}] = deployment.New(factory.KubernetesSharedInformerFactory())

	return &ResourceGetter{
		namespacedResourceGetters: namespacedResourceGetters,
		clusterResourceGetters:    clusterResourceGetters,
	}
}

// TryResource will retrieve a getter with resource name, it doesn't guarantee find resource with correct group version
// need to refactor this use schema.GroupVersionResource
func (r *ResourceGetter) TryResource(clusterScope bool, resource string) resources.Interface {
	if clusterScope {
		for k, v := range r.clusterResourceGetters {
			if k.Resource == resource {
				return v
			}
		}
	}
	for k, v := range r.namespacedResourceGetters {
		if k.Resource == resource {
			return v
		}
	}
	return nil
}

func (r *ResourceGetter) Get(resource, namespace, name string) (runtime.Object, error) {
	clusterScope := namespace == ""
	getter := r.TryResource(clusterScope, resource)
	if getter == nil {
		return nil, ErrResourceNotSupported
	}
	return getter.Get(namespace, name)
}

func (r *ResourceGetter) List(resource, namespace string, query *query.Query) (*resources.ListResult, error) {
	clusterScope := namespace == ""
	getter := r.TryResource(clusterScope, resource)
	if getter == nil {
		return nil, ErrResourceNotSupported
	}
	return getter.List(namespace, query)
}
