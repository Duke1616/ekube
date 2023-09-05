package impl

import (
	"context"
	v1 "ekube/api/pb/endpoint/v1"
	"ekube/internal/endpoint"
	"ekube/tools"
	"github.com/infraboard/mcube/types/ftime"
)

func (s *service) DeleteEndpoint(ctx context.Context, request *v1.DeleteEndpointRequest) (*v1.Endpoint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DescribeEndpoint(ctx context.Context, req *v1.DescribeEndpointRequest) (*v1.Endpoint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ListEndpoints(ctx context.Context, req *v1.ListEndpointRequest) (*v1.EndpointSet, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) RegistryEndpoint(ctx context.Context, req *v1.RegistryRequest) (*v1.RegistryResponse, error) {
	endpoints := make([]*v1.Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &v1.Endpoint{
			Id:        tools.GenHashID(req.ServiceId, req.Entries[i].Path),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: req.ServiceId,
			Entry:     req.Entries[i],
		}
		endpoints = append(endpoints, ep)
	}

	err := s.data.InsertMany(ctx, endpoints)
	if err != nil {
		return nil, err
	}

	return endpoint.NewRegistryResponse("ok"), nil
}
