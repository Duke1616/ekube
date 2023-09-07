package api

import (
	v1 "ekube/api/pb/workspace/v1"
	"ekube/internal/workspace"
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
	workspace workspace.Service
	log       logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(workspace.AppName)
	h.workspace = ioc.GetInternalApp(workspace.AppName).(workspace.Service)
	return nil
}

func (h *handler) Name() string {
	return workspace.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"企业空间"}

	ws.Route(ws.POST("/").To(h.CreateWorkspace).
		Doc("创建企业空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.CreateWorkspaceRequest{}).
		Writes(v1.Workspace{}))

	ws.Route(ws.GET("/").To(h.ListWorkspace).
		Doc("查询企业空间列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(v1.WorkspaceSet{}).
		Returns(200, "OK", v1.WorkspaceSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}").To(h.DescribeWorkspace).
		Doc("查询企业空间详情").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(v1.Workspace{}).
		Returns(200, "OK", v1.Workspace{}).
		Returns(404, "Not Found", nil))
}

func init() {
	ioc.RegistryRestfulApp(h)
}
