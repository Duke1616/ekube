package k8s

import (
	"ekube/pkg/apiserver/query"
	"k8s.io/apimachinery/pkg/runtime"
)

type Interface interface {
	// Get retrieves a single object by its namespace and name
	Get(namespace, name string) (runtime.Object, error)

	// List retrieves a collection of objects matches given query
	List(namespace string, query *query.Query) (*ListResult, error)
}

type ListResult struct {
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"total"`
}

func DefaultList(objects []runtime.Object, q *query.Query) *ListResult {
	return &ListResult{
		TotalItems: len(objects),
		Items:      objectsToInterfaces(objects),
	}
}

func objectsToInterfaces(objs []runtime.Object) []interface{} {
	res := make([]interface{}, 0)
	for _, obj := range objs {
		res = append(res, obj)
	}
	return res
}
