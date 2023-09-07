package user

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}
