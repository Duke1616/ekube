package impl

import (
	"context"
	v1 "ekube/api/pb/policy/v1"
	"ekube/internal/namespace"
	"ekube/internal/role"
	"ekube/tools"
)

func (s *service) CreatePolicy(ctx context.Context, req *v1.CreatePolicyRequest) (*v1.Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins := &v1.Policy{
		Meta: tools.NewMeta(),
		Spec: req,
	}
	ins.GenID()

	// 查看企业空间-项目是否存在
	if _, err := s.namespace.DescribeNamespace(ctx, namespace.NewDescribeNamespaceRequestByName(
		req.Namespace, req.Workspace)); err != nil {
		return nil, err
	}

	// 查看role是否存在
	if _, err := s.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(ins.Spec.RoleId)); err != nil {
		return nil, err
	}
	// 插入策略
	if err := s.data.Insert(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) ListPolicy(ctx context.Context, req *v1.ListPolicyRequest) (*v1.PolicySet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	set, err := s.data.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *service) DescribePolicy(ctx context.Context, req *v1.DescribePolicyRequest) (*v1.Policy, error) {
	return nil, nil
}
