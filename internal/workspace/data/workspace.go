package data

import (
	"context"
	v1 "ekube/api/pb/workspace/v1"
	"ekube/internal/workspace"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Workspace) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted workspace(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func (s *Data) Get(ctx context.Context, req *v1.DescribeWorkspaceRequest) (*v1.Workspace, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case v1.DESCRIBE_BY_NANE:
		filter["name"] = req.Name
	default:
		filter["_id"] = req.Id
	}

	ins := &v1.Workspace{}
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("workspace %s not found", req.Id)
		}

		return nil, exception.NewInternalServerError("find workspace %s error, %s", req.Id, err)
	}

	return ins, nil
}

func newListWorkspaceRequest(req *v1.ListWorkspaceRequest) *listWorkspaceRequest {
	return &listWorkspaceRequest{
		ListWorkspaceRequest: req}
}

type listWorkspaceRequest struct {
	*v1.ListWorkspaceRequest
}

func (r *listWorkspaceRequest) FindOptions() *options.FindOptions {
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

func (r *listWorkspaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.UserId != "" {
		filter["user_id"] = r.UserId
	}

	return filter
}

func (s *Data) List(ctx context.Context, req *v1.ListWorkspaceRequest) (*v1.WorkspaceSet, error) {
	query := newListWorkspaceRequest(req)
	s.log.Debugf("find filter: %s", query.FindFilter())

	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find workspace error, error is %s", err)
	}

	set := workspace.NewWorkspaceSet()
	for resp.Next(ctx) {
		ins := &v1.Workspace{}
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode workspace error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get workspace count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
