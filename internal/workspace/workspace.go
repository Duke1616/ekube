package workspace

import (
	v1 "ekube/api/pb/workspace/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
)

const (
	AppName = "workspace"
)

func NewListWorkspaceRequestFromHTTP(r *restful.Request) *v1.ListWorkspaceRequest {
	page := tools.NewPageRequestFromHTTP(r.Request)

	req := NewListWorkspaceRequest()
	req.Page = page

	return req
}

func NewListWorkspaceRequest() *v1.ListWorkspaceRequest {
	return &v1.ListWorkspaceRequest{
		Page: tools.NewDefaultPageRequest(),
	}
}
