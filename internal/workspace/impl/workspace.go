package impl

import (
	"context"
	workspaceV1 "ekube/api/pb/workspace/v1"
)

func (s *service) CreateWorkspace(ctx context.Context, req *workspaceV1.CreateWorkspaceRequest) (
	*workspaceV1.Workspace, error) {

	return nil, nil
}

func (s *service) UpdateWorkspace(ctx context.Context, req *workspaceV1.UpdateWorkspaceRequest) (
	*workspaceV1.Workspace, error) {
	return nil, nil
}

func (s *service) DeleteWorkspace(ctx context.Context, req *workspaceV1.DeleteWorkspaceRequest) (
	*workspaceV1.Workspace, error) {
	return nil, nil
}

func (s *service) DescribeWorkspace(ctx context.Context, req *workspaceV1.DescribeWorkspaceRequest) (
	*workspaceV1.Workspace, error) {
	return nil, nil
}

func (s *service) ListWorkspace(ctx context.Context, req *workspaceV1.ListWorkspaceRequest) (
	*workspaceV1.WorkspaceSet, error) {
	return nil, nil
}
