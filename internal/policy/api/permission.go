package api

import (
	"ekube/api/pb/role/v1"
	"ekube/internal/policy"
	"ekube/protocol/ioc"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type permissionHandler struct {
	policy policy.Service
	log    logger.Logger
}

func (h *permissionHandler) Config() error {
	h.log = zap.L().Named(policy.AppName)
	h.policy = ioc.GetInternalApp(policy.AppName).(policy.Service)
	return nil
}

func (h *permissionHandler) Name() string {
	return "permission"
}

func (h *permissionHandler) Version() string {
	return "v1"
}

func (h *permissionHandler) Registry(ws *restful.WebService) {
	tags := []string{"用户鉴权"}

	ws.Route(ws.POST("/").To(h.CheckPermission).
		Doc("权限校验").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Writes(role.Permission{}).
		Returns(200, "OK", role.Permission{}).
		Returns(404, "Not Found", nil))
}

func init() {
	ioc.RegistryRestfulApp(&permissionHandler{})
}
