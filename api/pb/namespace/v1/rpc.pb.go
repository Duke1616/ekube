// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/namespace/v1/rpc.proto

package namespace

import (
	page "ekube/third_party/ekube/pb/page"
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

type DESCRIBE_BY int32

const (
	// 通过Id
	DESCRIBE_BY_ID DESCRIBE_BY = 0
	// 通过名称
	DESCRIBE_BY_NANE DESCRIBE_BY = 1
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "ID",
		1: "NANE",
	}
	DESCRIBE_BY_value = map[string]int32{
		"ID":   0,
		"NANE": 1,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_namespace_v1_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_pb_namespace_v1_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_pb_namespace_v1_rpc_proto_rawDescGZIP(), []int{0}
}

type ListNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页
	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 根据用户权限过滤展示
	// @gotags: json:"cluster_id"
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id"`
	// 根据用户权限过滤展示
	// @gotags: json:"workspace"
	Workspace string `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace"`
}

func (x *ListNamespaceRequest) Reset() {
	*x = ListNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceRequest) ProtoMessage() {}

func (x *ListNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceRequest.ProtoReflect.Descriptor instead.
func (*ListNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ListNamespaceRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListNamespaceRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListNamespaceRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type DescribeNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 获取详情的方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=ekube.namespace.v1.DESCRIBE_BY" json:"describe_by"`
	// 集群Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 集群Id
	// @gotags: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 企业空间
	// @gotags: json:"workspace"
	Workspace string `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace"`
}

func (x *DescribeNamespaceRequest) Reset() {
	*x = DescribeNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeNamespaceRequest) ProtoMessage() {}

func (x *DescribeNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DescribeNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeNamespaceRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_ID
}

func (x *DescribeNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeNamespaceRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type UpdateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNamespaceRequest) Reset() {
	*x = UpdateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceRequest) ProtoMessage() {}

func (x *UpdateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_rpc_proto_rawDescGZIP(), []int{2}
}

type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 删除workspace
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pb_namespace_v1_rpc_proto protoreflect.FileDescriptor

var file_pb_namespace_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x9e, 0x01,
	0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59,
	0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x18,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x2a, 0x1f, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42,
	0x59, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4e,
	0x45, 0x10, 0x01, 0x32, 0xc4, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5b, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x65,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x60, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x2e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_namespace_v1_rpc_proto_rawDescOnce sync.Once
	file_pb_namespace_v1_rpc_proto_rawDescData = file_pb_namespace_v1_rpc_proto_rawDesc
)

func file_pb_namespace_v1_rpc_proto_rawDescGZIP() []byte {
	file_pb_namespace_v1_rpc_proto_rawDescOnce.Do(func() {
		file_pb_namespace_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_namespace_v1_rpc_proto_rawDescData)
	})
	return file_pb_namespace_v1_rpc_proto_rawDescData
}

var file_pb_namespace_v1_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_namespace_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_namespace_v1_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),                 // 0: ekube.namespace.v1.DESCRIBE_BY
	(*ListNamespaceRequest)(nil),     // 1: ekube.namespace.v1.ListNamespaceRequest
	(*DescribeNamespaceRequest)(nil), // 2: ekube.namespace.v1.DescribeNamespaceRequest
	(*UpdateNamespaceRequest)(nil),   // 3: ekube.namespace.v1.UpdateNamespaceRequest
	(*DeleteNamespaceRequest)(nil),   // 4: ekube.namespace.v1.DeleteNamespaceRequest
	(*page.PageRequest)(nil),         // 5: ekube.page.PageRequest
	(*NamespaceSet)(nil),             // 6: ekube.namespace.v1.NamespaceSet
	(*Namespace)(nil),                // 7: ekube.namespace.v1.Namespace
}
var file_pb_namespace_v1_rpc_proto_depIdxs = []int32{
	5, // 0: ekube.namespace.v1.ListNamespaceRequest.page:type_name -> ekube.page.PageRequest
	0, // 1: ekube.namespace.v1.DescribeNamespaceRequest.describe_by:type_name -> ekube.namespace.v1.DESCRIBE_BY
	1, // 2: ekube.namespace.v1.RPC.ListNamespace:input_type -> ekube.namespace.v1.ListNamespaceRequest
	2, // 3: ekube.namespace.v1.RPC.DescribeNamespace:input_type -> ekube.namespace.v1.DescribeNamespaceRequest
	6, // 4: ekube.namespace.v1.RPC.ListNamespace:output_type -> ekube.namespace.v1.NamespaceSet
	7, // 5: ekube.namespace.v1.RPC.DescribeNamespace:output_type -> ekube.namespace.v1.Namespace
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_namespace_v1_rpc_proto_init() }
func file_pb_namespace_v1_rpc_proto_init() {
	if File_pb_namespace_v1_rpc_proto != nil {
		return
	}
	file_pb_namespace_v1_namespace_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_namespace_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceRequest); i {
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
		file_pb_namespace_v1_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeNamespaceRequest); i {
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
		file_pb_namespace_v1_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceRequest); i {
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
		file_pb_namespace_v1_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
			RawDescriptor: file_pb_namespace_v1_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_namespace_v1_rpc_proto_goTypes,
		DependencyIndexes: file_pb_namespace_v1_rpc_proto_depIdxs,
		EnumInfos:         file_pb_namespace_v1_rpc_proto_enumTypes,
		MessageInfos:      file_pb_namespace_v1_rpc_proto_msgTypes,
	}.Build()
	File_pb_namespace_v1_rpc_proto = out.File
	file_pb_namespace_v1_rpc_proto_rawDesc = nil
	file_pb_namespace_v1_rpc_proto_goTypes = nil
	file_pb_namespace_v1_rpc_proto_depIdxs = nil
}
