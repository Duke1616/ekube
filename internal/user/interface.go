package user

import (
	"context"
	v1 "ekube/api/pb/user/v1"
)

const (
	AppName = "user"
)

type Service interface {
	v1.RPCServer
	CreateUser(context.Context, *v1.CreateUserRequest) (*v1.User, error)
}
