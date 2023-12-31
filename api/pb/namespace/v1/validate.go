package namespace

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreateNamespaceRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ListNamespaceRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribeNamespaceRequest) Validate() error {
	return validate.Struct(req)
}
