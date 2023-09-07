package impl

import (
	"context"
	v1 "ekube/api/pb/role/v1"
	"ekube/internal/role"
	"ekube/tools"
	"fmt"
)

func (s *service) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &v1.Role{
		Meta:        tools.NewMeta(),
		Spec:        req,
		Permissions: []*v1.PermissionSpec{},
	}

	r.Meta.Id = tools.FnvHash(r.FullName())

	// 保存角色信息
	err := s.data.Insert(ctx, r)
	if err != nil {
		return nil, err
	}

	// 录入相应权限
	addReq := role.NewAddPermissionToRoleRequest()
	addReq.CreateBy = req.CreateBy
	addReq.Permissions = req.Specs
	addReq.RoleId = r.Meta.Id

	fmt.Println(addReq.RoleId)
	return s.AddPermissionToRole(ctx, addReq)
}

func (s *service) DescribeRole(ctx context.Context, req *v1.DescribeRoleRequest) (*v1.Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins, err := s.data.GetRole(ctx, req)
	if err != nil {
		return nil, err
	}

	// 补充权限
	pReq := role.NewListPermissionRequest()
	pReq.RoleId = ins.Meta.Id
	pReq.Page = tools.NewPageRequest(role.MaxRolePermission, 1)
	ps, err := s.ListPermission(ctx, pReq)
	if err != nil {
		return nil, err
	}
	ins.Permissions = ps.PermissionSpecs()
	return ins, nil
}

func (s *service) ListRole(ctx context.Context, req *v1.ListRoleRequest) (*v1.RoleSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	set, err := s.data.ListRole(ctx, req)
	if err != nil {
		return nil, err
	}

	return set, nil
}
