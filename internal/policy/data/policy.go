package data

import (
	"context"
	v1 "ekube/api/pb/policy/v1"
	"ekube/internal/policy"
	"ekube/internal/role"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Policy) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Meta.Id, err)
	}

	return nil
}

func (s *Data) Get(ctx context.Context, req *v1.DescribePolicyRequest) (*v1.Policy, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case v1.DESCRIBE_BY_NANE:
		filter["name"] = req.Id
	default:
		filter["_id"] = req.Id
	}

	ins := &v1.Policy{}
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req.Id)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *Data) List(ctx context.Context, req *v1.ListPolicyRequest) (*v1.PolicySet, error) {
	query := newListPolicyRequest(req)
	s.log.Debugf("find filter: %s", query.FindFilter())

	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet()
	for resp.Next(ctx) {
		ins := &v1.Policy{}
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}

		// 补充关联的角色信息
		if req.WithRole {
			descRole := role.NewDescribeRoleRequestWithID(ins.Spec.RoleId)
			ins.Role, err = s.role.DescribeRole(ctx, descRole)
			if err != nil {
				return nil, err
			}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func newListPolicyRequest(req *v1.ListPolicyRequest) *listPolicyRequest {
	return &listPolicyRequest{
		ListPolicyRequest: req}
}

type listPolicyRequest struct {
	*v1.ListPolicyRequest
}

func (r *listPolicyRequest) FindOptions() *options.FindOptions {
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

func (r *listPolicyRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Workspace != "" {
		filter["workspace"] = r.Workspace
	}

	return filter
}
