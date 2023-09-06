package api

import (
	v1 "ekube/api/pb/namespace/v1"
	"ekube/internal/namespace"
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
	namespace namespace.Service
	log       logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(namespace.AppName)
	h.namespace = ioc.GetInternalApp(namespace.AppName).(namespace.Service)
	return nil
}

func (h *handler) Name() string {
	return namespace.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"集群项目"}

	ws.Route(ws.POST("/").To(h.CreateNamespace).
		Doc("创建项目").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.CreateNamespaceRequest{}).
		Writes(v1.Namespace{}))

	ws.Route(ws.GET("/").To(h.ListNamespace).
		Doc("查询项目列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(v1.NamespaceSet{}).
		Returns(200, "OK", v1.NamespaceSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	ioc.RegistryRestfulApp(h)
}
