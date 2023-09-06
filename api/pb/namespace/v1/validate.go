package namespace

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreateNamespaceRequest) Validate() error {
	return validate.Struct(req)
}
