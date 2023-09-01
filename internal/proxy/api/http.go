package api

import (
	"ekube/internal/cluster"
	"ekube/internal/proxy"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/resource"
	"ekube/protocol/ioc"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

var (
	Handler = &handler{}
)

type handler struct {
	service         cluster.ServiceCluster
	log             logger.Logger
	informerFactory informer.InformerFactory
	kubernetes      kubernetes.Interface
	resourceGetter  *resource.ResourceGetter
}

func (h *handler) AddToContainer(informerFactory informer.InformerFactory, kubernetes kubernetes.Interface) error {
	h.informerFactory = informerFactory
	h.kubernetes = kubernetes
	h.resourceGetter = resource.NewResourceGetter(h.informerFactory)
	return nil
}

func (h *handler) Config() error {
	h.log = zap.L().Named(proxy.AppName)
	h.service = ioc.GetInternalApp(cluster.AppName).(cluster.ServiceCluster)
	return nil
}

func (h *handler) Name() string {
	return proxy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	ws.Filter(h.ClusterMiddleware)
	h.registryPodHandler(ws)
}

func init() {
	// refer to https://github.com/kubernetes/kubernetes/issues/94688
	tools.RegisterConversions(scheme.Scheme)

	ioc.RegistryRestfulApp(Handler)
}

func (h *handler) ClusterMiddleware(
	req *restful.Request,
	resp *restful.Response,
	next *restful.FilterChain) {

	// 处理请求
	//clusterId := req.PathParameter("cluster_id")
	//if clusterId == "" {
	//	response.Failed(resp, fmt.Errorf("url path param cluster_id required"))
	//	return
	//}

	// 获取集群client对象
	//descReq := cluster.NewDescribeClusterRequest(clusterId)
	//ins, err := h.service.DescribeCluster(req.Request.Context(), descReq)
	//if err != nil {
	//	response.Failed(resp, fmt.Errorf("describe cluster_id error, %s", err))
	//	return
	//}

	//fmt.Println(ins)

	//client, err := resource.NewClient(ins.Spec.KubeConfig)
	//if err != nil {
	//	response.Failed(resp, fmt.Errorf("new k8s client error, %s", err))
	//	return
	//}
	//req.SetAttribute(proxy.ATTRIBUTE_K8S_CLIENT, client)

	// next flow
	next.ProcessFilter(req, resp)

	// 处理响应
}
