package role

import (
	"context"
	"ekube/api/pb/role/v1"
)

const (
	AppName = "role"
)

type Service interface {
	CreateRole(context.Context, *role.CreateRoleRequest) (*role.Role, error)
	AddPermissionToRole(context.Context, *role.AddPermissionToRoleRequest) (*role.Role, error)
	role.RPCServer
}
