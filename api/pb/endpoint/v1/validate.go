package endpoint

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}
