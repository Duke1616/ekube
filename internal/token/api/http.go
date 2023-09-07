package api

import (
	v1 "ekube/api/pb/token/v1"
	"ekube/internal/token"
	"ekube/protocol/ioc"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	token token.Service
	log   logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(token.AppName)
	h.token = ioc.GetInternalApp(token.AppName).(token.Service)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"登录管理"}

	ws.Route(ws.POST("/").To(h.IssueToken).
		Doc("颁发令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(v1.IssueTokenRequest{}).
		Writes(v1.Token{}).
		Returns(200, "OK", v1.Token{}))

	ws.Route(ws.POST("/validate").To(h.ValidateToken).
		Doc("验证令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(v1.ValidateTokenRequest{}).
		Writes(v1.Token{}).
		Returns(200, "OK", v1.Token{}))

}

func init() {
	ioc.RegistryRestfulApp(h)
}
