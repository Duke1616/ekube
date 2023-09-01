package workspace

import (
	"context"
	workspaceV1 "ekube/api/pb/workspace/v1"
)

type Service interface {
	ServiceWorkspace
}

type ServiceWorkspace interface {
	workspaceV1.RPCServer
	CreateWorkspace(context.Context, *workspaceV1.CreateWorkspaceRequest) (*workspaceV1.Workspace, error)
	UpdateWorkspace(context.Context, *workspaceV1.UpdateWorkspaceRequest) (*workspaceV1.Workspace, error)
	DeleteWorkspace(context.Context, *workspaceV1.DeleteWorkspaceRequest) (*workspaceV1.Workspace, error)
}
