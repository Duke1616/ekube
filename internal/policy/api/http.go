package api

import (
	v1 "ekube/api/pb/policy/v1"
	"ekube/internal/policy"
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
	policy policy.Service
	log    logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(policy.AppName)
	h.policy = ioc.GetInternalApp(policy.AppName).(policy.Service)
	return nil
}

func (h *handler) Name() string {
	return policy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"策略管理"}

	ws.Route(ws.POST("/").To(h.CreatePolicy).
		Doc("创建策略").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.CreatePolicyRequest{}).
		Writes(v1.Policy{}))

	ws.Route(ws.GET("/").To(h.ListPolicy).
		Doc("查询策略列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(v1.PolicySet{}).
		Returns(200, "OK", v1.PolicySet{}).
		Returns(404, "Not Found", nil))

}

func init() {
	ioc.RegistryRestfulApp(h)
}
