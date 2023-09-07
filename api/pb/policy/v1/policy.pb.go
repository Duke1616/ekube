// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.2
// source: pb/policy/v1/policy.proto

package policy

import (
	v1 "ekube/api/pb/role/v1"
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

// Policy 权限策略
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 策略定义
	// @gotags: bson:",inline" json:"spec"
	Spec *CreatePolicyRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 关联的角色对象
	// @gotags: bson:"-" json:"role,omitempty"
	Role *v1.Role `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty" bson:"-"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_policy_v1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_pb_policy_v1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_pb_policy_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Policy) GetSpec() *CreatePolicyRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Policy) GetRole() *v1.Role {
	if x != nil {
		return x.Role
	}
	return nil
}

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建者
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 策略所属企业空间
	// @gotags: bson:"workspace" json:"workspace"
	Workspace string `protobuf:"bytes,2,opt,name=workspace,proto3" json:"workspace" bson:"workspace"`
	// 所属项目
	// @gotags: bson:"namespace" json:"namespace" validate:"lte=120"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" bson:"namespace" validate:"lte=120"`
	// 用户Id
	// @gotags: bson:"user_id" json:"user_id" validate:"required,lte=120"
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id" validate:"required,lte=120"`
	// 角色Id
	// @gotags: bson:"role_id" json:"role_id" validate:"required,lte=40"
	RoleId string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id" validate:"required,lte=40"`
	// 该角色的生效范围
	// @gotags: bson:"scope" json:"scope"
	Scope string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope" bson:"scope"`
	// 策略过期时间
	// @gotags: bson:"expired_time" json:"expired_time"
	ExpiredTime int64 `protobuf:"varint,7,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time" bson:"expired_time"`
	// 启用该策略
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 扩展属性
	// @gotags: bson:"extra" json:"extra"
	Extra map[string]string `protobuf:"bytes,10,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_policy_v1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_policy_v1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pb_policy_v1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreatePolicyRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *CreatePolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreatePolicyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreatePolicyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *CreatePolicyRequest) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *CreatePolicyRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreatePolicyRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *CreatePolicyRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// PolicySet
type PolicySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Policy `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *PolicySet) Reset() {
	*x = PolicySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_policy_v1_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySet) ProtoMessage() {}

func (x *PolicySet) ProtoReflect() protoreflect.Message {
	mi := &file_pb_policy_v1_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySet.ProtoReflect.Descriptor instead.
func (*PolicySet) Descriptor() ([]byte, []int) {
	return file_pb_policy_v1_policy_proto_rawDescGZIP(), []int{2}
}

func (x *PolicySet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PolicySet) GetItems() []*Policy {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pb_policy_v1_policy_proto protoreflect.FileDescriptor

var file_pb_policy_v1_policy_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01,
	0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0xf9, 0x03, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x48,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a,
	0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42,
	0x1f, 0x5a, 0x1d, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_policy_v1_policy_proto_rawDescOnce sync.Once
	file_pb_policy_v1_policy_proto_rawDescData = file_pb_policy_v1_policy_proto_rawDesc
)

func file_pb_policy_v1_policy_proto_rawDescGZIP() []byte {
	file_pb_policy_v1_policy_proto_rawDescOnce.Do(func() {
		file_pb_policy_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_policy_v1_policy_proto_rawDescData)
	})
	return file_pb_policy_v1_policy_proto_rawDescData
}

var file_pb_policy_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_policy_v1_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),              // 0: ekube.policy.v1.Policy
	(*CreatePolicyRequest)(nil), // 1: ekube.policy.v1.CreatePolicyRequest
	(*PolicySet)(nil),           // 2: ekube.policy.v1.PolicySet
	nil,                         // 3: ekube.policy.v1.CreatePolicyRequest.ExtraEntry
	nil,                         // 4: ekube.policy.v1.CreatePolicyRequest.LabelsEntry
	(*meta.Meta)(nil),           // 5: ekube.meta.Meta
	(*v1.Role)(nil),             // 6: ekube.role.v1.Role
}
var file_pb_policy_v1_policy_proto_depIdxs = []int32{
	5, // 0: ekube.policy.v1.Policy.meta:type_name -> ekube.meta.Meta
	1, // 1: ekube.policy.v1.Policy.spec:type_name -> ekube.policy.v1.CreatePolicyRequest
	6, // 2: ekube.policy.v1.Policy.role:type_name -> ekube.role.v1.Role
	3, // 3: ekube.policy.v1.CreatePolicyRequest.extra:type_name -> ekube.policy.v1.CreatePolicyRequest.ExtraEntry
	4, // 4: ekube.policy.v1.CreatePolicyRequest.labels:type_name -> ekube.policy.v1.CreatePolicyRequest.LabelsEntry
	0, // 5: ekube.policy.v1.PolicySet.items:type_name -> ekube.policy.v1.Policy
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_policy_v1_policy_proto_init() }
func file_pb_policy_v1_policy_proto_init() {
	if File_pb_policy_v1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_policy_v1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_pb_policy_v1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_pb_policy_v1_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySet); i {
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
			RawDescriptor: file_pb_policy_v1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_policy_v1_policy_proto_goTypes,
		DependencyIndexes: file_pb_policy_v1_policy_proto_depIdxs,
		MessageInfos:      file_pb_policy_v1_policy_proto_msgTypes,
	}.Build()
	File_pb_policy_v1_policy_proto = out.File
	file_pb_policy_v1_policy_proto_rawDesc = nil
	file_pb_policy_v1_policy_proto_goTypes = nil
	file_pb_policy_v1_policy_proto_depIdxs = nil
}
