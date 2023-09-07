package data

import (
	"context"
	v1 "ekube/api/pb/user/v1"
	"ekube/internal/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) Insert(ctx context.Context, ins *v1.User) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			ins.Spec.Username, err)
	}

	return nil
}
func (s *Data) Get(ctx context.Context, req *v1.DescribeUserRequest) (*v1.User, error) {
	filter := bson.M{}

	filter["username"] = req.Username

	resp := &v1.User{}
	if err := s.col.FindOne(ctx, filter).Decode(resp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("user %s error, %s", req, err)
	}

	return resp, nil
}

func (s *Data) List(ctx context.Context, req *v1.ListUserRequest) (*v1.UserSet, error) {
	r := newListRequest(req)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	set := user.NewUserSet()
	// 循环
	for resp.Next(ctx) {
		ins := &v1.User{}
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode user error, error is %s", err)
		}
		if ins.Password != nil {
			ins.Password.Password = ""
			ins.Password.History = []string{}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get user count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func newListRequest(r *v1.ListUserRequest) *listRequest {
	return &listRequest{
		r,
	}
}

type listRequest struct {
	*v1.ListUserRequest
}

func (r *listRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *listRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
