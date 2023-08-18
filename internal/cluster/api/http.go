package api

import (
	v1 "ekube/api/pb/cluster/v1"
	"ekube/internal/cluster"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/http/label"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service cluster.ClusterService
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(cluster.AppName)
	h.service = ioc.GetInternalApp(cluster.AppName).(cluster.ClusterService)
	return nil
}

func (h *handler) Name() string {
	return cluster.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"集群管理"}

	ws.Route(ws.POST("/").To(h.CreateCluster).
		Doc("创建集群").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.CreateClusterRequest{}).
		Writes(v1.Cluster{}))

	ws.Route(ws.GET("/").To(h.ListCluster).
		Doc("查看集群").
		Metadata(label.Resource, cluster.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("vendor", "厂商").DataType("string")).
		Param(ws.QueryParameter("region", "地域").DataType("string")).
		Writes(v1.ClusterSet{}).
		Returns(200, "OK", v1.ClusterSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}").To(h.DescribeCluster).
		Doc("集群详情").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(v1.Cluster{}).
		Returns(200, "OK", v1.Cluster{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteCluster).
		Doc("删除集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))

	ws.Route(ws.PUT("/{id}").To(h.PutCluster).
		Doc("修改集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(v1.Cluster{}).
		Returns(200, "OK", v1.Cluster{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchCluster).
		Doc("修改集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(v1.Cluster{}).
		Returns(200, "OK", v1.Cluster{}).
		Returns(404, "Not Found", nil))
}

func init() {
	ioc.RegistryRestfulApp(h)
}
