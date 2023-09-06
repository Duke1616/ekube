package api

import (
	v1 "ekube/api/pb/role/v1"
	"ekube/internal/role"
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
	role role.Service
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(role.AppName)
	h.role = ioc.GetInternalApp(role.AppName).(role.Service)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"角色管理"}

	ws.Route(ws.POST("/").To(h.CreateRole).
		Doc("创建角色").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.CreateRoleRequest{}).
		Writes(v1.Role{}))

	ws.Route(ws.GET("/").To(h.ListRole).
		Doc("查询角色列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(v1.RoleSet{}).
		Returns(200, "OK", v1.RoleSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	ioc.RegistryRestfulApp(h)
}
