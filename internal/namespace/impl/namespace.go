package impl

import (
	"context"
	v1 "ekube/api/pb/namespace/v1"
	"ekube/internal/cluster"
	"ekube/internal/workspace"
	"ekube/tools"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreateNamespace(ctx context.Context, req *v1.CreateNamespaceRequest) (*v1.Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// 检验集群是否存在
	if _, err := s.cluster.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.ClusterId)); err != nil {
		return nil, err
	}

	// 检查workspace是否存在
	if _, err := s.workspace.DescribeWorkspace(ctx, workspace.NewDescribeWorkspaceRequestByName(
		req.Workspace)); err != nil {
		return nil, err
	}

	ins := &v1.Namespace{
		Meta: tools.NewMeta(),
		Spec: req,
	}

	ins.Meta.Id = tools.FnvHash(req.Name, req.Workspace, req.ClusterId)

	err := s.data.Insert(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) ListNamespace(ctx context.Context, req *v1.ListNamespaceRequest) (*v1.NamespaceSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	set, err := s.data.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *service) DescribeNamespace(ctx context.Context, req *v1.DescribeNamespaceRequest) (*v1.Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins, err := s.data.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
