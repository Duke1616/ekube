package role

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *AddPermissionToRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribeRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ListPermissionRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ListRoleRequest) Validate() error {
	return validate.Struct(req)
}
