package impl

import (
	"context"
	v1 "ekube/api/pb/role/v1"
	"ekube/internal/role"
	"ekube/tools"
	"github.com/infraboard/mcube/exception"
)

func (s *service) AddPermissionToRole(ctx context.Context, req *v1.AddPermissionToRoleRequest) (*v1.Role, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add permission error, %s", err)
	}

	ins, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.RoleId))
	if err != nil {
		return nil, err
	}

	// 查询角色条目数是否超标
	queryPerm := role.NewListPermissionRequest()
	queryPerm.Page = tools.NewPageRequest(role.MaxRolePermission, 1)
	queryPerm.SkipItems = true
	queryPerm.RoleId = ins.Meta.Id

	ps, err := s.ListPermission(ctx, queryPerm)
	if err != nil {
		return nil, err
	}

	if ps.Total+int64(req.Length()) > role.MaxRolePermission {
		return nil, exception.NewBadRequest("一个角色最多可以添加%d权限条目, 当前条目数: %d, 新增条目数: %d",
			role.MaxRolePermission, ps.Total, req.Length())
	}

	perms := role.NewPermission(ins.Meta.Id, req.Permissions...)

	insMany := role.TransferPermissionToDocs(perms)
	if _, err = s.data.InsertPermission(ctx, insMany); err != nil {
		return nil, exception.NewInternalServerError("inserted permission(%s) document error, %s",
			perms, err)
	}

	return s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.RoleId))
}

func (s *service) ListPermission(ctx context.Context, req *v1.ListPermissionRequest) (*v1.PermissionSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	set, err := s.data.ListPermission(ctx, req)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *service) DescribePermission(ctx context.Context, req *v1.DescribePermissionRequest) (*v1.Permission, error) {
	//TODO implement me
	panic("implement me")
}
