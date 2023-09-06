package impl

import (
	"context"
	v1 "ekube/api/pb/workspace/v1"
	"ekube/tools"
)

func (s *service) CreateWorkspace(ctx context.Context, req *v1.CreateWorkspaceRequest) (
	*v1.Workspace, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &v1.Workspace{
		Meta: tools.NewMeta(),
		Spec: req,
	}

	err := s.data.Insert(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service) UpdateWorkspace(ctx context.Context, req *v1.UpdateWorkspaceRequest) (
	*v1.Workspace, error) {
	return nil, nil
}

func (s *service) DeleteWorkspace(ctx context.Context, req *v1.DeleteWorkspaceRequest) (
	*v1.Workspace, error) {
	return nil, nil
}

func (s *service) DescribeWorkspace(ctx context.Context, req *v1.DescribeWorkspaceRequest) (
	*v1.Workspace, error) {
	return nil, nil
}

func (s *service) ListWorkspace(ctx context.Context, req *v1.ListWorkspaceRequest) (
	*v1.WorkspaceSet, error) {
	return nil, nil
}
