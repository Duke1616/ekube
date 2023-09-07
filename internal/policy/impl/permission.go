package impl

import (
	"context"
	policyV1 "ekube/api/pb/policy/v1"
	roleV1 "ekube/api/pb/role/v1"
	"ekube/internal/endpoint"
	"ekube/internal/policy"
	"ekube/internal/role"
	"ekube/tools"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CheckPermission(ctx context.Context, req *policyV1.CheckPermissionRequest) (
	*roleV1.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 查询用户权限策略
	pReq := policy.NewListPolicyRequest()
	pReq.Username = req.Username
	pReq.Workspace = req.Workspace
	pReq.Namespace = req.Namespace
	pReq.WithRole = true
	ps, err := s.ListPolicy(ctx, pReq)
	if err != nil {
		return nil, err
	}

	// 无用户相关权限策略设置
	if len(ps.Items) == 0 {
		return nil, exception.NewPermissionDeny("no permission")
	}

	// 查询用户需要鉴权的功能
	fn := endpoint.NewDescribeEndpointRequestWithID(tools.GenHashID(req.ServiceId, req.Path))
	ep, err := s.endpoint.DescribeEndpoint(ctx, fn)
	if err != nil {
		return nil, err
	}

	// 判断改功能是否需要鉴权
	if !ep.Entry.PermissionEnable {
		return role.NewSkipPermission("endpoint not enable permission check, allow all access"), nil
	}

	// 判断策略是否允许
	var perm *roleV1.Permission
	for i := range ps.Items {
		p := ps.Items[i]
		permOk, ok, err := p.Role.HasPermission(ep)
		if err != nil {
			return nil, err
		}
		if ok {
			perm = role.NewPermissionFromSpec(p.Spec.RoleId, permOk)
			perm.Scope = p.Spec.Scope
			s.log.Debugf("check roles %s has permission access endpoint [%s]", p.Role.Spec.Name, ep.Entry)
		}
	}

	if perm == nil {
		return nil, exception.NewPermissionDeny("in namespace %s, role %s has no permission access endpoint: %s",
			req.Namespace,
			ps.RoleNames(),
			ep.Entry.Path,
		)
	}

	return perm, nil
}
