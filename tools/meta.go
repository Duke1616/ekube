package tools

import (
	"ekube/third_party/ekube/pb/meta"
	"github.com/rs/xid"
	"time"
)

func NewMeta() *meta.Meta {
	return &meta.Meta{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
	}
}
