package api

import (
	"context"
	"io"
	"net/http"

	"ekube/utils/response"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/util/flushwriter"
	"k8s.io/client-go/kubernetes/scheme"

	clusterv1 "ekube/api/pb/cluster/v1"
	"ekube/pkg/apiserver/query"
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
		Reads(clusterv1.ListClusterRequest{}).
		Writes(corev1.PodList{}).
		Returns(200, "OK", corev1.PodList{}))

	ws.Route(ws.GET("/namespaces/{namespace}/pods/{pod}/log").To(h.PodLog).
		Doc("查看pod日志").
		Param(ws.PathParameter("namespace", "the watching namespace of the gateway")).
		Param(ws.PathParameter("pod", "the pod name of the gateway")).
		Returns(http.StatusOK, "OK", corev1.PodLogOptions{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func (h *handler) ListPods(r *restful.Request, w *restful.Response) {
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

func (h *handler) PodLog(r *restful.Request, w *restful.Response) {
	namespace := r.PathParameter("namespace")
	pod := r.PathParameter("pod")

	urlQuery := r.Request.URL.Query()
	logOptions := &corev1.PodLogOptions{}
	if err := scheme.ParameterCodec.DecodeParameters(urlQuery, corev1.SchemeGroupVersion, logOptions); err != nil {
		response.Failed(w, err)
		return
	}

	fw := flushwriter.Wrap(w.ResponseWriter)
	err := h.getPodLogs(r.Request.Context(), namespace, pod, logOptions, fw)
	if err != nil {
		response.Failed(w, err)
		return
	}
}

func (h *handler) getPodLogs(ctx context.Context, namespace string, podName string, logOptions *corev1.PodLogOptions,
	responseWriter io.Writer) error {
	podLogRequest := h.kubernetes.CoreV1().
		Pods(namespace).
		GetLogs(podName, logOptions)
	reader, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	_, err = io.Copy(responseWriter, reader)
	if err != nil {
		return err
	}
	return nil
}
