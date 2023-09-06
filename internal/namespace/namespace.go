package namespace

import (
	v1 "ekube/api/pb/namespace/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
)

func NewListNamespaceRequestFromHTTP(r *restful.Request) *v1.ListNamespaceRequest {
	page := tools.NewPageRequestFromHTTP(r.Request)

	req := NewListNamespaceRequest()
	req.Page = page

	return req
}

func NewListNamespaceRequest() *v1.ListNamespaceRequest {
	return &v1.ListNamespaceRequest{
		Page: tools.NewDefaultPageRequest(),
	}
}
