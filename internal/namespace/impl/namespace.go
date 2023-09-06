package impl

import (
	"context"
	v1 "ekube/api/pb/namespace/v1"
	"ekube/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateNamespace(ctx context.Context, req *v1.CreateNamespaceRequest) (*v1.Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &v1.Namespace{
		Meta: tools.NewMeta(),
		Spec: req,
	}

	err := s.data.Insert(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service) ListNamespace(ctx context.Context, req *v1.ListNamespaceRequest) (*v1.NamespaceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespace not implemented")
}

func (s *service) DescribeNamespace(ctx context.Context, req *v1.DescribeNamespaceRequest) (*v1.Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNamespace not implemented")
}
