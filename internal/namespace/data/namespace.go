package data

import (
	"context"
	v1 "ekube/api/pb/namespace/v1"
	"github.com/infraboard/mcube/exception"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Namespace) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted workspace(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}
