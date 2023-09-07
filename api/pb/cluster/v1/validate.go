package cluster

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (x *CreateClusterRequest) Validate() error {
	return validate.Struct(x)
}

func (x *Cluster) IsAlive() error {
	if x.Status == nil {
		return fmt.Errorf("status is nil")
	}

	//if !x.Status.IsAlive {
	//	return fmt.Errorf(x.Status.Message)
	//}

	return nil
}
