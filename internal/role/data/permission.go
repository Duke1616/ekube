package data

import (
	"context"
	v1 "ekube/api/pb/role/v1"
	"ekube/internal/role"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) InsertPermission(ctx context.Context, ins []interface{}) (*mongo.InsertManyResult, error) {
	if _, err := s.perm.InsertMany(ctx, ins); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Data) ListPermission(ctx context.Context, req *v1.ListPermissionRequest) (*v1.PermissionSet, error) {
	query, err := newListPermissionRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query permission filter: %s", query.FindFilter())
	resp, err := s.perm.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find permissionn error, error is %s", err)
	}

	// 循环
	set := role.NewPermissionSet()
	if !req.SkipItems {
		for resp.Next(ctx) {
			ins := &v1.Permission{}
			if err = resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode permission error, error is %s", err)
			}
			set.Add(ins)
		}
	}

	count, err := s.perm.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get permission count error, error is %s", err)
	}

	set.Total = count

	return set, nil
}

func newListPermissionRequest(req *v1.ListPermissionRequest) (*listPermissionRequest, error) {
	return &listPermissionRequest{
		ListPermissionRequest: req}, nil
}

type listPermissionRequest struct {
	*v1.ListPermissionRequest
}

func (r *listPermissionRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *listPermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}

	return filter
}
