package api

import (
	v1 "ekube/api/pb/endpoint/v1"
	"ekube/internal/endpoint"
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
	service endpoint.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(endpoint.AppName)
	h.service = ioc.GetInternalApp(endpoint.AppName).(endpoint.Service)
	return nil
}

func (h *handler) Name() string {
	return endpoint.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务功能"}
	ws.Route(ws.POST("/").To(h.RegistryEndpoint).
		Doc("注册服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(v1.RegistryRequest{}).
		Writes(v1.EndpointSet{}))

	ws.Route(ws.GET("/").To(h.ListEndpoints).
		Doc("查询服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Reads(v1.ListEndpointRequest{}).
		Writes(v1.EndpointSet{}).
		Returns(200, "OK", endpoint.NewEndpointSet()))

	ws.Route(ws.GET("/{id}").To(h.DescribeEndpoint).
		Doc("查询服务功能详情").
		Param(ws.PathParameter("id", "identifier of the service").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Writes(v1.Endpoint{}).
		Returns(200, "OK", v1.Endpoint{}).
		Returns(404, "Not Found", nil))

}

func init() {
	ioc.RegistryRestfulApp(h)
}
