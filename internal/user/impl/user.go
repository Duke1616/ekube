package impl

import (
	"context"
	v1 "ekube/api/pb/user/v1"
	"ekube/internal/user"
	"ekube/tools"
	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
)

func (s *service) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	pass, err := user.NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u := &v1.User{
		Meta:     tools.NewMeta(),
		Spec:     req,
		Password: pass,
		Status:   &v1.Status{},
	}

	u.Meta.Id = xid.New().String()

	if err = s.data.Insert(ctx, u); err != nil {
		return nil, err
	}

	u.Password = nil
	return u, nil
}

func (s *service) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.UserSet, error) {
	return s.data.List(ctx, req)
}

func (s *service) DescribeUser(ctx context.Context, req *v1.DescribeUserRequest) (*v1.User, error) {
	return s.data.Get(ctx, req)
}
