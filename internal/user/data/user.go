package data

import (
	"context"
	v1 "ekube/api/pb/user/v1"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
