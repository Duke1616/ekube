package role

import (
	v1 "ekube/api/pb/role/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
)

func NewCreateRoleRequest() *v1.CreateRoleRequest {
	return &v1.CreateRoleRequest{
		Labels:  map[string]string{},
		Specs:   []*v1.PermissionSpec{},
		Enabled: true,
	}
}

func NewListPermissionRequest() *v1.ListPermissionRequest {
	return &v1.ListPermissionRequest{
		Page: tools.NewDefaultPageRequest(),
	}
}

func NewListRoleRequestFromHTTP(r *restful.Request) *v1.ListRoleRequest {
	page := tools.NewPageRequestFromHTTP(r.Request)

	req := NewListRoleRequest()
	req.Page = page
	req.WithPermission = r.QueryParameter("with_permission") == "true"
	return req
}

func NewListRoleRequest() *v1.ListRoleRequest {
	return &v1.ListRoleRequest{
		Page: tools.NewDefaultPageRequest(),
	}
}

func NewRoleSet() *v1.RoleSet {
	return &v1.RoleSet{
		Items: []*v1.Role{},
	}
}

func NewDefaultRole() *v1.Role {
	return &v1.Role{
		Meta:        tools.NewMeta(),
		Spec:        NewCreateRoleRequest(),
		Permissions: []*v1.PermissionSpec{},
	}
}
