package api

import (
	v1 "ekube/api/pb/user/v1"
	"ekube/internal/user"
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
	user user.Service
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.user = ioc.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"账号管理"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("创建账号").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(v1.TYPE_PRIMARY)).
		Reads(v1.CreateUserRequest{}).
		Returns(200, "创建成功", &v1.User{}))

	ws.Route(ws.GET("/").To(h.ListUser).
		Doc("查询子账号列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(v1.TYPE_PRIMARY)).
		Returns(200, "OK", v1.UserSet{}))
}

func init() {
	ioc.RegistryRestfulApp(h)
}
