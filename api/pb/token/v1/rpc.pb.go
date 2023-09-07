// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/token/v1/rpc.proto

package token

import (
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

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_token_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_token_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_pb_token_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_pb_token_v1_rpc_proto protoreflect.FileDescriptor

var file_pb_token_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x39, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x53, 0x0a, 0x03, 0x52,
	0x50, 0x43, 0x12, 0x4c, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_token_v1_rpc_proto_rawDescOnce sync.Once
	file_pb_token_v1_rpc_proto_rawDescData = file_pb_token_v1_rpc_proto_rawDesc
)

func file_pb_token_v1_rpc_proto_rawDescGZIP() []byte {
	file_pb_token_v1_rpc_proto_rawDescOnce.Do(func() {
		file_pb_token_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_token_v1_rpc_proto_rawDescData)
	})
	return file_pb_token_v1_rpc_proto_rawDescData
}

var file_pb_token_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_token_v1_rpc_proto_goTypes = []interface{}{
	(*ValidateTokenRequest)(nil), // 0: ekube.token.v1.ValidateTokenRequest
	(*Token)(nil),                // 1: ekube.token.v1.Token
}
var file_pb_token_v1_rpc_proto_depIdxs = []int32{
	0, // 0: ekube.token.v1.RPC.ValidateToken:input_type -> ekube.token.v1.ValidateTokenRequest
	1, // 1: ekube.token.v1.RPC.ValidateToken:output_type -> ekube.token.v1.Token
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_token_v1_rpc_proto_init() }
func file_pb_token_v1_rpc_proto_init() {
	if File_pb_token_v1_rpc_proto != nil {
		return
	}
	file_pb_token_v1_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_token_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
			RawDescriptor: file_pb_token_v1_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_token_v1_rpc_proto_goTypes,
		DependencyIndexes: file_pb_token_v1_rpc_proto_depIdxs,
		MessageInfos:      file_pb_token_v1_rpc_proto_msgTypes,
	}.Build()
	File_pb_token_v1_rpc_proto = out.File
	file_pb_token_v1_rpc_proto_rawDesc = nil
	file_pb_token_v1_rpc_proto_goTypes = nil
	file_pb_token_v1_rpc_proto_depIdxs = nil
}