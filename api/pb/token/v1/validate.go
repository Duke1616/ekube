package token

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func (req *ValidateTokenRequest) Validate() error {
	return validate.Struct(req)
}
