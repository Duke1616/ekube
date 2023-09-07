// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.2
// source: pb/policy/v1/rpc.proto

package policy

import (
	context "context"
	v1 "ekube/api/pb/role/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPC_ListPolicy_FullMethodName      = "/ekube.policy.v1.RPC/ListPolicy"
	RPC_DescribePolicy_FullMethodName  = "/ekube.policy.v1.RPC/DescribePolicy"
	RPC_CheckPermission_FullMethodName = "/ekube.policy.v1.RPC/CheckPermission"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 查询企业空间列表
	ListPolicy(ctx context.Context, in *ListPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error)
	// 查询企业空间详情
	DescribePolicy(ctx context.Context, in *DescribePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// 策略鉴权
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*v1.Permission, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) ListPolicy(ctx context.Context, in *ListPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error) {
	out := new(PolicySet)
	err := c.cc.Invoke(ctx, RPC_ListPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribePolicy(ctx context.Context, in *DescribePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, RPC_DescribePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*v1.Permission, error) {
	out := new(v1.Permission)
	err := c.cc.Invoke(ctx, RPC_CheckPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 查询企业空间列表
	ListPolicy(context.Context, *ListPolicyRequest) (*PolicySet, error)
	// 查询企业空间详情
	DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error)
	// 策略鉴权
	CheckPermission(context.Context, *CheckPermissionRequest) (*v1.Permission, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) ListPolicy(context.Context, *ListPolicyRequest) (*PolicySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicy not implemented")
}
func (UnimplementedRPCServer) DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePolicy not implemented")
}
func (UnimplementedRPCServer) CheckPermission(context.Context, *CheckPermissionRequest) (*v1.Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_ListPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).ListPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_ListPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).ListPolicy(ctx, req.(*ListPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribePolicy(ctx, req.(*DescribePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ekube.policy.v1.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPolicy",
			Handler:    _RPC_ListPolicy_Handler,
		},
		{
			MethodName: "DescribePolicy",
			Handler:    _RPC_DescribePolicy_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _RPC_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/policy/v1/rpc.proto",
}