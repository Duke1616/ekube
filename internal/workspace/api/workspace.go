package api

import (
	v1 "ekube/api/pb/workspace/v1"
	"ekube/internal/workspace"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateWorkspace(r *restful.Request, w *restful.Response) {
	req := &v1.CreateWorkspaceRequest{}

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.workspace.CreateWorkspace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListWorkspace(r *restful.Request, w *restful.Response) {
	req := workspace.NewListWorkspaceRequestFromHTTP(r)
	set, err := h.workspace.ListWorkspace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeWorkspace(r *restful.Request, w *restful.Response) {
	req := workspace.NewDescribeWorkspaceRequestById(r.PathParameter("id"))
	set, err := h.workspace.DescribeWorkspace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
