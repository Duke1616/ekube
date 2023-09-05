package endpoint

import (
	v1 "ekube/api/pb/endpoint/v1"
	"github.com/infraboard/mcube/http/request"
	"net/http"
	"strings"
)

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *v1.RegistryRequest {
	return &v1.RegistryRequest{
		Entries: []*v1.Entry{},
	}
}

// NewQueryEndpointRequestFromHTTP 列表查询请求
func NewQueryEndpointRequestFromHTTP(r *http.Request) *v1.ListEndpointRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	query := &v1.ListEndpointRequest{
		Page:         page,
		Path:         qs.Get("path"),
		Method:       qs.Get("method"),
		FunctionName: qs.Get("function_name"),
	}

	sids := qs.Get("service_ids")
	if sids != "" {
		query.ServiceIds = strings.Split(sids, ",")
	}
	rs := qs.Get("resources")
	if rs != "" {
		query.Resources = strings.Split(rs, ",")
	}

	return query
}

func NewDescribeEndpointRequestWithID(id string) *v1.DescribeEndpointRequest {
	return &v1.DescribeEndpointRequest{Id: id}
}

// NewEndpointSet 实例化
func NewEndpointSet() *v1.EndpointSet {
	return &v1.EndpointSet{
		Items: []*v1.Endpoint{},
	}
}

func NewRegistryResponse(ok string) *v1.RegistryResponse {
	return &v1.RegistryResponse{
		Message: ok,
	}
}
