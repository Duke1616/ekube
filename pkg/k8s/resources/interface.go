package resources

import (
	"ekube/pkg/apiserver/query"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sort"
	"strings"
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

// CompareFunc 对比进行
type CompareFunc func(runtime.Object, runtime.Object, query.Field) bool

// FilterFunc 根据标签等过滤
type FilterFunc func(runtime.Object, query.Filter) bool

type TransformFunc func(runtime.Object) runtime.Object

type ListOptions struct {
	CompareFunc    CompareFunc
	FilterFunc     FilterFunc
	TransformFuncs []TransformFunc
}

func DefaultList(objects []runtime.Object, q *query.Query, options ListOptions) *ListResult {
	var filtered []runtime.Object

	for _, object := range objects {
		selected := true
		for field, value := range q.Filters {
			if !options.FilterFunc(object, query.Filter{Field: field, Value: value}) {
				selected = false
				break
			}
		}

		if selected {
			for _, transform := range options.TransformFuncs {
				object = transform(object)
			}
			filtered = append(filtered, object)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		if !q.Ascending {
			return options.CompareFunc(filtered[i], filtered[j], q.SortBy)
		}
		return !options.CompareFunc(filtered[i], filtered[j], q.SortBy)
	})

	total := len(filtered)

	if q.Pagination == nil {
		q.Pagination = query.NoPagination
	}

	start, end := q.Pagination.GetValidPagination(total)

	return &ListResult{
		TotalItems: len(filtered),
		Items:      objectsToInterfaces(filtered[start:end]),
	}
}

func DefaultObjectMetaCompare(left, right metav1.ObjectMeta, sortBy query.Field) bool {
	switch sortBy {
	// ?sortBy=name
	case query.FieldName:
		return strings.Compare(left.Name, right.Name) > 0
	//	?sortBy=creationTimestamp
	default:
		fallthrough
	case query.FieldCreateTime:
		fallthrough
	case query.FieldCreationTimeStamp:
		// compare by name if creation timestamp is equal
		if left.CreationTimestamp.Equal(&right.CreationTimestamp) {
			return strings.Compare(left.Name, right.Name) > 0
		}
		return left.CreationTimestamp.After(right.CreationTimestamp.Time)
	}
}

func objectsToInterfaces(objs []runtime.Object) []interface{} {
	res := make([]interface{}, 0)
	for _, obj := range objs {
		res = append(res, obj)
	}
	return res
}
