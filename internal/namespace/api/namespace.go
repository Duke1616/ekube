package api

import (
	v1 "ekube/api/pb/namespace/v1"
	"ekube/internal/namespace"
	"ekube/utils/response"
	"github.com/emicklei/go-restful/v3"
)

func (h *handler) CreateNamespace(r *restful.Request, w *restful.Response) {
	req := &v1.CreateNamespaceRequest{}

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.namespace.CreateNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewListNamespaceRequestFromHTTP(r)
	set, err := h.namespace.ListNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
