package namespace

import (
	"context"
	"ekube/api/pb/namespace/v1"
)

const (
	AppName = "namespace"
)

type Service interface {
	namespace.RPCServer
	CreateNamespace(context.Context, *namespace.CreateNamespaceRequest) (*namespace.Namespace, error)
}
