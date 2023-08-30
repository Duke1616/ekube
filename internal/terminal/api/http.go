package api

import (
	"ekube/pkg/k8s/terminal"
	"ekube/protocol/ioc"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	Handler = &handler{}
)

type handler struct {
	client kubernetes.Interface
	log    logger.Logger
	config *rest.Config
	option *terminal.Option
}

func (h *handler) AddToContainer(client kubernetes.Interface, config *rest.Config, option *terminal.Option) error {
	h.config = config
	h.client = client
	h.option = option
	return nil
}

func (h *handler) Config() error {
	h.log = zap.L().Named("terminal")
	return nil
}

func (h *handler) Name() string {
	return "terminal"
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	h.registryPodHandler(ws)
}

func init() {
	ioc.RegistryRestfulApp(Handler)
}

func (h *handler) registryPodHandler(ws *restful.WebService) {
	tags := []string{"[terminal] 终端管理"}

	termHandler := newTerminalHandler(h.client, h.config, h.option)

	ws.Route(ws.GET("/namespaces/{namespace}/pods/{pod}/exec").
		To(termHandler.handleTerminalSession).
		Param(ws.PathParameter("namespace", "namespace of which the pod located in")).
		Param(ws.PathParameter("pod", "name of the pod")).
		Doc("create terminal session").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}
