package registry

import (
	"context"
	"io"
)

type Registry interface {
	Registry(ctx context.Context, si ServiceInstance) error
	UnRegistry(ctx context.Context, si ServiceInstance) error
	ListServices(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	Subscribe(serviceName string) (<-chan Event, error)

	io.Closer
}

type ServiceInstance struct {
	Name string
	// 定位信息
	Address string
	// 权重
	Weight uint32
	// 分组字断
	Group string
}

type Event struct {
}
