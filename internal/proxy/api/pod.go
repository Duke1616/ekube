package api

import (
	v1 "ekube/api/pb/cluster/v1"
	"ekube/pkg/apiserver/query"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryPodHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] Pod管理"}

	ws.Route(ws.GET("/{cluster_id}/pods").To(h.ListPods).
		Doc("查询Pod列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.ListClusterRequest{}).
		Writes(corev1.PodList{}).
		Returns(200, "OK", corev1.PodList{}))
}

func (h *handler) ListPods(r *restful.Request, w *restful.Response) {
	//client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*resource.KubernetesClient)
	//
	//req := meta.NewListRequestFromHttp(r.Request)
	////ins, err := client.Pod().ListPod(r.Request.Context(), req)
	//if err != nil {
	//	response.Failed(w, err)
	//	return
	//}

	q := query.ParseQueryParameter(r)
	resourceType := r.QueryParameter("resources")
	namespace := r.QueryParameter("namespace")

	result, err := h.resourceGetter.List(resourceType, namespace, q)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, result)
}
