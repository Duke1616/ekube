package impl

import (
	"context"
	v1 "ekube/api/pb/workspace/v1"
	"ekube/internal/cluster"
	"ekube/tools"
)

func (s *service) CreateWorkspace(ctx context.Context, req *v1.CreateWorkspaceRequest) (
	*v1.Workspace, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 检验集群是否存在
	_, err := s.cluster.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.ClusterId[0]))
	if err != nil {
		return nil, err
	}

	ins := &v1.Workspace{
		Meta: tools.NewMeta(),
		Spec: req,
	}

	ins.Meta.Id = tools.FnvHash(req.Name)

	err = s.data.Insert(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeWorkspace(ctx context.Context, req *v1.DescribeWorkspaceRequest) (
	*v1.Workspace, error) {
	ins, err := s.data.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) ListWorkspace(ctx context.Context, req *v1.ListWorkspaceRequest) (
	*v1.WorkspaceSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	set, err := s.data.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *service) UpdateWorkspace(ctx context.Context, req *v1.UpdateWorkspaceRequest) (
	*v1.Workspace, error) {
	return nil, nil
}

func (s *service) DeleteWorkspace(ctx context.Context, req *v1.DeleteWorkspaceRequest) (
	*v1.Workspace, error) {
	return nil, nil
}
