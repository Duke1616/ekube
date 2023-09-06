package api

import (
	"ekube/internal/role"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateRole(r *restful.Request, w *restful.Response) {
	req := role.NewCreateRoleRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.role.CreateRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListRole(r *restful.Request, w *restful.Response) {
	req := role.NewListRoleRequestFromHTTP(r)
	set, err := h.role.ListRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
