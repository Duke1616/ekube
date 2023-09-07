package data

import (
	"context"
	v1 "ekube/api/pb/namespace/v1"
	"ekube/internal/namespace"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Namespace) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted namespace(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func (s *Data) Get(ctx context.Context, req *v1.DescribeNamespaceRequest) (*v1.Namespace, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case v1.DESCRIBE_BY_NANE:
		filter["name"] = req.Name
		filter["workspace"] = req.Workspace
	default:
		filter["_id"] = req.Id
	}

	ins := &v1.Namespace{}
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("namespace ID: %s, Name: %s not found", req.Id, req.Name)
		}

		return nil, exception.NewInternalServerError("find namespace %s error, %s", req.Id, err)
	}

	return ins, nil
}

func newListNamespaceRequest(req *v1.ListNamespaceRequest) *listNamespaceRequest {
	return &listNamespaceRequest{
		ListNamespaceRequest: req}
}

type listNamespaceRequest struct {
	*v1.ListNamespaceRequest
}

func (r *listNamespaceRequest) FindOptions() *options.FindOptions {
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

func (r *listNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ClusterId != "" {
		filter["cluster_id"] = r.ClusterId
	}

	if r.Workspace != "" {
		filter["workspace"] = r.Workspace
	}

	return filter
}

func (s *Data) List(ctx context.Context, req *v1.ListNamespaceRequest) (*v1.NamespaceSet, error) {
	query := newListNamespaceRequest(req)
	s.log.Debugf("find filter: %s", query.FindFilter())

	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find namespace error, error is %s", err)
	}

	set := namespace.NewNamespaceSet()
	for resp.Next(ctx) {
		ins := &v1.Namespace{}
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode namespace error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get namespace count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
