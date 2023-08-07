package api

import (
	"ekube/internal/cluster"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) ListCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewListClusterRequestFromHTTP(r)

	ins, err := h.service.ListCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewCreateClusterRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewDescribeClusterRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Desensitization()
	response.Success(w, ins)
}

func (h *handler) DeleteCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewDeleteClusterRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutCluster(r *restful.Request, w *restful.Response) {
	//tk := r.Attribute("token").(*token.Token)

	req := cluster.NewPutClusterRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	//req.UpdateBy = tk.Username

	set, err := h.service.UpdateCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchCluster(r *restful.Request, w *restful.Response) {
	//tk := r.Attribute("token").(*token.Token)
	req := cluster.NewPatchClusterRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	//req.UpdateBy = tk.Username

	set, err := h.service.UpdateCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
