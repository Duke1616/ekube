// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/cluster/v1/rpc.proto

package cluster

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
	resource "github.com/infraboard/mcube/pb/resource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// 资源标签过滤
	// @gotags: json:"filters"
	Filters []*resource.LabelRequirement `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters"`
	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page"`
	// 关键字参数
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	// 集群所属厂商
	// @gotags: json:"vendor"
	Vendor string `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor"`
	// 集群所属地域
	// @gotags: json:"region"
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region"`
}

func (x *ListClusterRequest) Reset() {
	*x = ListClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cluster_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClusterRequest) ProtoMessage() {}

func (x *ListClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cluster_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClusterRequest.ProtoReflect.Descriptor instead.
func (*ListClusterRequest) Descriptor() ([]byte, []int) {
	return file_pb_cluster_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ListClusterRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *ListClusterRequest) GetFilters() []*resource.LabelRequirement {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListClusterRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListClusterRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ListClusterRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *ListClusterRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type DescribeClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 集群Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeClusterRequest) Reset() {
	*x = DescribeClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cluster_v1_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeClusterRequest) ProtoMessage() {}

func (x *DescribeClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cluster_v1_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeClusterRequest.ProtoReflect.Descriptor instead.
func (*DescribeClusterRequest) Descriptor() ([]byte, []int) {
	return file_pb_cluster_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 更新时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 更新的书本信息
	// @gotags: json:"spec"
	Spec *CreateClusterRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec"`
}

func (x *UpdateClusterRequest) Reset() {
	*x = UpdateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cluster_v1_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClusterRequest) ProtoMessage() {}

func (x *UpdateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cluster_v1_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) {
	return file_pb_cluster_v1_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateClusterRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateClusterRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateClusterRequest) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UpdateClusterRequest) GetSpec() *CreateClusterRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cluster_v1_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cluster_v1_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_pb_cluster_v1_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pb_cluster_v1_rpc_proto protoreflect.FileDescriptor

var file_pb_cluster_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x45,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x3a, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xb0, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x51, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0f,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x28, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_cluster_v1_rpc_proto_rawDescOnce sync.Once
	file_pb_cluster_v1_rpc_proto_rawDescData = file_pb_cluster_v1_rpc_proto_rawDesc
)

func file_pb_cluster_v1_rpc_proto_rawDescGZIP() []byte {
	file_pb_cluster_v1_rpc_proto_rawDescOnce.Do(func() {
		file_pb_cluster_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_cluster_v1_rpc_proto_rawDescData)
	})
	return file_pb_cluster_v1_rpc_proto_rawDescData
}

var file_pb_cluster_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_cluster_v1_rpc_proto_goTypes = []interface{}{
	(*ListClusterRequest)(nil),        // 0: ekube.cluster.v1.ListClusterRequest
	(*DescribeClusterRequest)(nil),    // 1: ekube.cluster.v1.DescribeClusterRequest
	(*UpdateClusterRequest)(nil),      // 2: ekube.cluster.v1.UpdateClusterRequest
	(*DeleteClusterRequest)(nil),      // 3: ekube.cluster.v1.DeleteClusterRequest
	(*resource.Scope)(nil),            // 4: infraboard.mcube.resource.Scope
	(*resource.LabelRequirement)(nil), // 5: infraboard.mcube.resource.LabelRequirement
	(*request.PageRequest)(nil),       // 6: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),          // 7: infraboard.mcube.request.UpdateMode
	(*CreateClusterRequest)(nil),      // 8: ekube.cluster.v1.CreateClusterRequest
	(*ClusterSet)(nil),                // 9: ekube.cluster.v1.ClusterSet
	(*Cluster)(nil),                   // 10: ekube.cluster.v1.Cluster
}
var file_pb_cluster_v1_rpc_proto_depIdxs = []int32{
	4,  // 0: ekube.cluster.v1.ListClusterRequest.scope:type_name -> infraboard.mcube.resource.Scope
	5,  // 1: ekube.cluster.v1.ListClusterRequest.filters:type_name -> infraboard.mcube.resource.LabelRequirement
	6,  // 2: ekube.cluster.v1.ListClusterRequest.page:type_name -> infraboard.mcube.page.PageRequest
	7,  // 3: ekube.cluster.v1.UpdateClusterRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	8,  // 4: ekube.cluster.v1.UpdateClusterRequest.spec:type_name -> ekube.cluster.v1.CreateClusterRequest
	0,  // 5: ekube.cluster.v1.RPC.ListCluster:input_type -> ekube.cluster.v1.ListClusterRequest
	1,  // 6: ekube.cluster.v1.RPC.DescribeCluster:input_type -> ekube.cluster.v1.DescribeClusterRequest
	9,  // 7: ekube.cluster.v1.RPC.ListCluster:output_type -> ekube.cluster.v1.ClusterSet
	10, // 8: ekube.cluster.v1.RPC.DescribeCluster:output_type -> ekube.cluster.v1.Cluster
	7,  // [7:9] is the sub-list for method output_type
	5,  // [5:7] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_cluster_v1_rpc_proto_init() }
func file_pb_cluster_v1_rpc_proto_init() {
	if File_pb_cluster_v1_rpc_proto != nil {
		return
	}
	file_pb_cluster_v1_cluster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_cluster_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cluster_v1_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cluster_v1_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cluster_v1_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_cluster_v1_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_cluster_v1_rpc_proto_goTypes,
		DependencyIndexes: file_pb_cluster_v1_rpc_proto_depIdxs,
		MessageInfos:      file_pb_cluster_v1_rpc_proto_msgTypes,
	}.Build()
	File_pb_cluster_v1_rpc_proto = out.File
	file_pb_cluster_v1_rpc_proto_rawDesc = nil
	file_pb_cluster_v1_rpc_proto_goTypes = nil
	file_pb_cluster_v1_rpc_proto_depIdxs = nil
}
