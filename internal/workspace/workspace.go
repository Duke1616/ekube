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

func NewWorkspaceSet() *v1.WorkspaceSet {
	return &v1.WorkspaceSet{
		Items: []*v1.Workspace{},
	}
}

func NewDescribeWorkspaceRequestByName(name string) *v1.DescribeWorkspaceRequest {
	return &v1.DescribeWorkspaceRequest{
		DescribeBy: v1.DESCRIBE_BY_NANE,
		Name:       name,
	}
}

func NewDescribeWorkspaceRequestById(id string) *v1.DescribeWorkspaceRequest {
	return &v1.DescribeWorkspaceRequest{
		DescribeBy: v1.DESCRIBE_BY_ID,
		Id:         id,
	}
}
