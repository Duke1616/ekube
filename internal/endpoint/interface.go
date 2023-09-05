package endpoint

import (
	"context"
	"ekube/api/pb/endpoint/v1"
)

const (
	AppName = "endpoint"
)

type Service interface {
	endpoint.RPCServer
	DeleteEndpoint(context.Context, *endpoint.DeleteEndpointRequest) (*endpoint.Endpoint, error)
}
