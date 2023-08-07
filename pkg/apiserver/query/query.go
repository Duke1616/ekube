package query

import (
	"github.com/emicklei/go-restful/v3"
	"k8s.io/apimachinery/pkg/labels"
	"strconv"
)

type Query struct {
	Pagination *Pagination

	// sort result in which field, default to FieldCreationTimeStamp
	SortBy Field

	// sort result in ascending or descending order, default to descending
	Ascending bool

	Filters map[Field]Value

	LabelSelector string
}

type Filter struct {
	Field Field
	Value Value
}

type Pagination struct {
	// items per page
	Limit int

	// offset
	Offset int
}

func (q *Query) Selector() labels.Selector {
	if selector, err := labels.Parse(q.LabelSelector); err != nil {
		return labels.Everything()
	} else {
		return selector
	}
}

var NoPagination = newPagination(-1, 0)

func newPagination(limit int, offset int) *Pagination {
	return &Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

func (p *Pagination) GetValidPagination(total int) (startIndex, endIndex int) {

	// no pagination
	if p.Limit == NoPagination.Limit {
		return 0, total
	}

	// out of range
	if p.Limit < 0 || p.Offset < 0 || p.Offset > total {
		return 0, 0
	}

	startIndex = p.Offset
	endIndex = startIndex + p.Limit

	if endIndex > total {
		endIndex = total
	}

	return startIndex, endIndex
}

func New() *Query {
	return &Query{
		Pagination: NoPagination,
		SortBy:     "",
		Ascending:  false,
		Filters:    map[Field]Value{},
	}
}

func ParseQueryParameter(request *restful.Request) *Query {
	query := New()

	limit, err := strconv.Atoi(request.QueryParameter(ParameterLimit))
	// equivalent to undefined, use the default value
	if err != nil {
		limit = -1
	}
	page, err := strconv.Atoi(request.QueryParameter(ParameterPage))
	// equivalent to undefined, use the default value
	if err != nil {
		page = 1
	}

	query.Pagination = newPagination(limit, (page-1)*limit)

	query.SortBy = Field(defaultString(request.QueryParameter(ParameterOrderBy), FieldCreationTimeStamp))

	ascending, err := strconv.ParseBool(defaultString(request.QueryParameter(ParameterAscending), "false"))
	if err != nil {
		query.Ascending = false
	} else {
		query.Ascending = ascending
	}

	query.LabelSelector = request.QueryParameter(ParameterLabelSelector)

	for key, values := range request.Request.URL.Query() {
		if !HasString([]string{ParameterPage, ParameterLimit, ParameterOrderBy, ParameterAscending, ParameterLabelSelector}, key) {
			// support multiple query condition
			for _, value := range values {
				query.Filters[Field(key)] = Value(value)
			}
		}
	}

	return query
}

func defaultString(value, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func HasString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
