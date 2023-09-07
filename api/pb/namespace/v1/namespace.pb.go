// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/namespace/v1/namespace.proto

package namespace

import (
	meta "ekube/third_party/ekube/pb/meta"
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

// Namespace 名称空间-项目绑定
type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 基础信息
	// @gotags: json:"spec" bson:",inline"
	Spec *CreateNamespaceRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Namespace) GetSpec() *CreateNamespaceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项目名称
	// @gotags: json:"name" bson:"name" validate:"required,lte=64"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=64"`
	// 项目别名 TODO 对应集群namespace
	// @gotags: json:"alias" bson:"alias" validate:"required,lte=64"
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias" bson:"alias" validate:"required,lte=64"`
	// 项目管理员
	// @gotags: json:"owner" bson:"owner"
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" bson:"owner"`
	// 项目绑定企业空间
	// @gotags: json:"workspace" bson:"workspace" validate:"required,lte=64"
	Workspace string `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace" bson:"workspace" validate:"required,lte=64"`
	// 项目绑定集群
	// @gotags: json:"cluster_id" bson:"cluster_id" validate:"required,lte=64"
	ClusterId string `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id" bson:"cluster_id" validate:"required,lte=64"`
	// 项目描述
	// @gotags: json:"description" bson:"description"
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description" bson:"description"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreateNamespaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateNamespaceRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *CreateNamespaceRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateNamespaceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// NamespaceSet
type NamespaceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Namespace `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *NamespaceSet) Reset() {
	*x = NamespaceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_namespace_v1_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceSet) ProtoMessage() {}

func (x *NamespaceSet) ProtoReflect() protoreflect.Message {
	mi := &file_pb_namespace_v1_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceSet.ProtoReflect.Descriptor instead.
func (*NamespaceSet) Descriptor() ([]byte, []int) {
	return file_pb_namespace_v1_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NamespaceSet) GetItems() []*Namespace {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pb_namespace_v1_namespace_proto protoreflect.FileDescriptor

var file_pb_namespace_v1_namespace_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0xb7, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0c,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x65, 0x6b, 0x75, 0x62, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_namespace_v1_namespace_proto_rawDescOnce sync.Once
	file_pb_namespace_v1_namespace_proto_rawDescData = file_pb_namespace_v1_namespace_proto_rawDesc
)

func file_pb_namespace_v1_namespace_proto_rawDescGZIP() []byte {
	file_pb_namespace_v1_namespace_proto_rawDescOnce.Do(func() {
		file_pb_namespace_v1_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_namespace_v1_namespace_proto_rawDescData)
	})
	return file_pb_namespace_v1_namespace_proto_rawDescData
}

var file_pb_namespace_v1_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_namespace_v1_namespace_proto_goTypes = []interface{}{
	(*Namespace)(nil),              // 0: ekube.namespace.v1.Namespace
	(*CreateNamespaceRequest)(nil), // 1: ekube.namespace.v1.CreateNamespaceRequest
	(*NamespaceSet)(nil),           // 2: ekube.namespace.v1.NamespaceSet
	(*meta.Meta)(nil),              // 3: ekube.meta.Meta
}
var file_pb_namespace_v1_namespace_proto_depIdxs = []int32{
	3, // 0: ekube.namespace.v1.Namespace.meta:type_name -> ekube.meta.Meta
	1, // 1: ekube.namespace.v1.Namespace.spec:type_name -> ekube.namespace.v1.CreateNamespaceRequest
	0, // 2: ekube.namespace.v1.NamespaceSet.items:type_name -> ekube.namespace.v1.Namespace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_namespace_v1_namespace_proto_init() }
func file_pb_namespace_v1_namespace_proto_init() {
	if File_pb_namespace_v1_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_namespace_v1_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_pb_namespace_v1_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_pb_namespace_v1_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceSet); i {
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
			RawDescriptor: file_pb_namespace_v1_namespace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_namespace_v1_namespace_proto_goTypes,
		DependencyIndexes: file_pb_namespace_v1_namespace_proto_depIdxs,
		MessageInfos:      file_pb_namespace_v1_namespace_proto_msgTypes,
	}.Build()
	File_pb_namespace_v1_namespace_proto = out.File
	file_pb_namespace_v1_namespace_proto_rawDesc = nil
	file_pb_namespace_v1_namespace_proto_goTypes = nil
	file_pb_namespace_v1_namespace_proto_depIdxs = nil
}
