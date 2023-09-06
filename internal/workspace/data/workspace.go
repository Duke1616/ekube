package data

import (
	"context"
	v1 "ekube/api/pb/workspace/v1"
	"github.com/infraboard/mcube/exception"
)

func (s *Data) Insert(ctx context.Context, ins *v1.Workspace) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted workspace(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}
