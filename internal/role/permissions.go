package role

import (
	v1 "ekube/api/pb/role/v1"
	"time"
)

const (
	// MaxRolePermission 一个角色允许添加的最大权限数
	MaxRolePermission = 500
)

func NewAddPermissionToRoleRequest() *v1.AddPermissionToRoleRequest {
	return &v1.AddPermissionToRoleRequest{
		Permissions: []*v1.PermissionSpec{},
	}
}

func NewDescribeRoleRequestWithID(id string) *v1.DescribeRoleRequest {
	return &v1.DescribeRoleRequest{
		Id: id,
	}
}

func NewPermission(roleId string, perms ...*v1.PermissionSpec) []*v1.Permission {
	var set []*v1.Permission
	for i := range perms {
		set = append(set, NewPermissionFromSpec(roleId, perms[i]))
	}
	return set
}

func NewPermissionFromSpec(roleId string, spec *v1.PermissionSpec) *v1.Permission {
	return &v1.Permission{
		Id:       spec.HashID(roleId),
		CreateAt: time.Now().Unix(),
		RoleId:   roleId,
		Spec:     spec,
	}
}

func TransferPermissionToDocs(perms []*v1.Permission) []interface{} {
	var docs []interface{}
	for i := range perms {
		docs = append(docs, perms[i])
	}
	return docs
}

func NewPermissionSet() *v1.PermissionSet {
	return &v1.PermissionSet{
		Items: []*v1.Permission{},
	}
}

// NewSkipPermission 如果endpoint不需要鉴权-触发
func NewSkipPermission(message string) *v1.Permission {
	return &v1.Permission{
		Spec: &v1.PermissionSpec{
			Effect: v1.EffectType_ALLOW,
			Desc:   message,
		},
	}
}
