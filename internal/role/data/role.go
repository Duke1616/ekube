package data

import (
	"context"
	v1 "ekube/api/pb/role/v1"
	"ekube/internal/role"
	"ekube/tools"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Role) error {
	if _, err := s.role.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted role(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func (s *Data) GetRole(ctx context.Context, req *v1.DescribeRoleRequest) (*v1.Role, error) {
	query, err := newDescribeRoleRequest(req)
	if err != nil {
		return nil, err
	}

	ins := &v1.Role{
		Meta:        tools.NewMeta(),
		Spec:        role.NewCreateRoleRequest(),
		Permissions: []*v1.PermissionSpec{},
	}

	if err = s.role.FindOne(ctx, query.FindFilter(), query.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("role %s not found", req)
		}

		return nil, exception.NewInternalServerError("find role %s error, %s", req, err)
	}

	return ins, nil
}

func (s *Data) ListRole(ctx context.Context, req *v1.ListRoleRequest) (*v1.RoleSet, error) {
	query, err := newListRoleRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query role filter: %s", query.FindFilter())
	resp, err := s.role.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := role.NewRoleSet()
	// 循环
	for resp.Next(ctx) {
		ins := role.NewDefaultRole()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode role error, error is %s", err)
		}
		// 补充权限
		if req.WithPermission {
			pReq := role.NewListPermissionRequest()
			pReq.RoleId = ins.Meta.Id
			pReq.Page = tools.NewPageRequest(role.MaxRolePermission, 1)
			ps, err := s.ListPermission(ctx, pReq)
			if err != nil {
				return nil, err
			}
			ins.Permissions = ps.PermissionSpecs()
		}
		set.Add(ins)
	}

	// count
	count, err := s.role.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get token count error, error is %s", err)
	}

	set.Total = count

	return set, nil
}

func newDescribeRoleRequest(req *v1.DescribeRoleRequest) (*describeRoleRequest, error) {
	return &describeRoleRequest{req}, nil
}

type describeRoleRequest struct {
	*v1.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

func (req *describeRoleRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newListRoleRequest(req *v1.ListRoleRequest) (*listRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &listRoleRequest{
		ListRoleRequest: req}, nil
}

type listRoleRequest struct {
	*v1.ListRoleRequest
}

func (r *listRoleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *listRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
