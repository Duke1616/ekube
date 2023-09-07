package workspace

import (
	"context"
	v1 "ekube/api/pb/workspace/v1"
)

type Service interface {
	v1.RPCServer
	CreateWorkspace(context.Context, *v1.CreateWorkspaceRequest) (*v1.Workspace, error)
	UpdateWorkspace(context.Context, *v1.UpdateWorkspaceRequest) (*v1.Workspace, error)
	DeleteWorkspace(context.Context, *v1.DeleteWorkspaceRequest) (*v1.Workspace, error)
}
