package policy

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CheckPermissionRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ListPolicyRequest) Validate() error {
	return validate.Struct(req)
}
