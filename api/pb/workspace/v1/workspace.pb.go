// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/workspace/v1/workspace.proto

package workspace

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

// Workspace 企业空间相关信息
type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 基础信息
	// @gotags: json:"spec" bson:",inline"
	Spec *CreateWorkspaceRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_workspace_v1_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_pb_workspace_v1_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_pb_workspace_v1_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *Workspace) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Workspace) GetSpec() *CreateWorkspaceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 企业空间名称
	// @gotags: json:"name" bson:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 空间管理员
	// @gotags: json:"owner" bson:"owner"
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner" bson:"owner"`
	// 空间描述信息
	// @gotags: json:"cluster_id" bson:"cluster_id"
	ClusterId []string `protobuf:"bytes,3,rep,name=cluster_id,json=clusterId,proto3" json:"cluster_id" bson:"cluster_id"`
	// 空间描述信息
	// @gotags: json:"description" bson:"description"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" bson:"description"`
}

func (x *CreateWorkspaceRequest) Reset() {
	*x = CreateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_workspace_v1_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest) ProtoMessage() {}

func (x *CreateWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_workspace_v1_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_pb_workspace_v1_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkspaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWorkspaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateWorkspaceRequest) GetClusterId() []string {
	if x != nil {
		return x.ClusterId
	}
	return nil
}

func (x *CreateWorkspaceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// WorkspaceSet
type WorkspaceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Workspace `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *WorkspaceSet) Reset() {
	*x = WorkspaceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_workspace_v1_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSet) ProtoMessage() {}

func (x *WorkspaceSet) ProtoReflect() protoreflect.Message {
	mi := &file_pb_workspace_v1_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSet.ProtoReflect.Descriptor instead.
func (*WorkspaceSet) Descriptor() ([]byte, []int) {
	return file_pb_workspace_v1_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *WorkspaceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *WorkspaceSet) GetItems() []*Workspace {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pb_workspace_v1_workspace_proto protoreflect.FileDescriptor

var file_pb_workspace_v1_workspace_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x62, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x62, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_workspace_v1_workspace_proto_rawDescOnce sync.Once
	file_pb_workspace_v1_workspace_proto_rawDescData = file_pb_workspace_v1_workspace_proto_rawDesc
)

func file_pb_workspace_v1_workspace_proto_rawDescGZIP() []byte {
	file_pb_workspace_v1_workspace_proto_rawDescOnce.Do(func() {
		file_pb_workspace_v1_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_workspace_v1_workspace_proto_rawDescData)
	})
	return file_pb_workspace_v1_workspace_proto_rawDescData
}

var file_pb_workspace_v1_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_workspace_v1_workspace_proto_goTypes = []interface{}{
	(*Workspace)(nil),              // 0: ekube.workspace.v1.Workspace
	(*CreateWorkspaceRequest)(nil), // 1: ekube.workspace.v1.CreateWorkspaceRequest
	(*WorkspaceSet)(nil),           // 2: ekube.workspace.v1.WorkspaceSet
	(*meta.Meta)(nil),              // 3: ekube.meta.Meta
}
var file_pb_workspace_v1_workspace_proto_depIdxs = []int32{
	3, // 0: ekube.workspace.v1.Workspace.meta:type_name -> ekube.meta.Meta
	1, // 1: ekube.workspace.v1.Workspace.spec:type_name -> ekube.workspace.v1.CreateWorkspaceRequest
	0, // 2: ekube.workspace.v1.WorkspaceSet.items:type_name -> ekube.workspace.v1.Workspace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_workspace_v1_workspace_proto_init() }
func file_pb_workspace_v1_workspace_proto_init() {
	if File_pb_workspace_v1_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_workspace_v1_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
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
		file_pb_workspace_v1_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceRequest); i {
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
		file_pb_workspace_v1_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSet); i {
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
			RawDescriptor: file_pb_workspace_v1_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_workspace_v1_workspace_proto_goTypes,
		DependencyIndexes: file_pb_workspace_v1_workspace_proto_depIdxs,
		MessageInfos:      file_pb_workspace_v1_workspace_proto_msgTypes,
	}.Build()
	File_pb_workspace_v1_workspace_proto = out.File
	file_pb_workspace_v1_workspace_proto_rawDesc = nil
	file_pb_workspace_v1_workspace_proto_goTypes = nil
	file_pb_workspace_v1_workspace_proto_depIdxs = nil
}
