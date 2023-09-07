package workspace

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreateWorkspaceRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ListWorkspaceRequest) Validate() error {
	return validate.Struct(req)
}
