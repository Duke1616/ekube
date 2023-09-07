package policy

import (
	v1 "ekube/api/pb/policy/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
)

func NewListPolicyRequestFromHTTP(r *restful.Request) *v1.ListPolicyRequest {
	page := tools.NewPageRequestFromHTTP(r.Request)

	req := NewListPolicyRequest()
	req.Page = page

	return req
}

func NewListPolicyRequest() *v1.ListPolicyRequest {
	return &v1.ListPolicyRequest{
		Page:          tools.NewDefaultPageRequest(),
		WithRole:      false,
		WithNamespace: false,
	}
}

func NewCheckPermissionRequest() *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Page: tools.NewPageRequest(100, 1),
	}
}

func NewPolicySet() *v1.PolicySet {
	return &v1.PolicySet{
		Items: []*v1.Policy{},
	}
}
