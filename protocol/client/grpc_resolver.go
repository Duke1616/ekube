package client

import (
	"context"
	"ekube/registry"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"time"
)

type grpcResolverBuilder struct {
	r       registry.Registry
	timeout time.Duration
}

func NewRegistryBuilder(r registry.Registry, timeout time.Duration) (*grpcResolverBuilder, error) {
	return &grpcResolverBuilder{
		r:       r,
		timeout: timeout,
	}, nil
}

func (b *grpcResolverBuilder) Scheme() string {
	return "registry"
}

func (b *grpcResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &grpcResolver{
		cc:      cc,
		r:       b.r,
		target:  target,
		timeout: b.timeout,
	}

	r.resolve()
	go r.watch()
	return r, nil
}

type grpcResolver struct {
	target  resolver.Target
	r       registry.Registry
	cc      resolver.ClientConn
	timeout time.Duration
	close   chan struct{}
}

func (g *grpcResolver) ResolveNow(options resolver.ResolveNowOptions) {
	g.resolve()
}

func (g *grpcResolver) watch() {
	events, err := g.r.Subscribe(g.target.Endpoint())
	if err != nil {
		g.cc.ReportError(err)
		return
	}
	select {
	case <-events:
		g.ResolveNow(resolver.ResolveNowOptions{})
	case <-g.close:
		return

	}
}

func (g *grpcResolver) resolve() {
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	instances, err := g.r.ListServices(ctx, g.target.Endpoint())
	if err != nil {
		g.cc.ReportError(err)
		return
	}

	address := make([]resolver.Address, 0, len(instances))
	for _, si := range instances {
		address = append(address, resolver.Address{
			Addr: si.Address,
			Attributes: attributes.New("weight", si.Weight).
				WithValue("group", si.Group),
		})
	}
	err = g.cc.UpdateState(resolver.State{
		Addresses: address,
	})

	if err != nil {
		g.cc.ReportError(err)
		return
	}
}

func (g *grpcResolver) Close() {
	close(g.close)
}
