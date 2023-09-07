package api

import (
	v1 "ekube/api/pb/policy/v1"
	"ekube/internal/policy"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreatePolicy(r *restful.Request, w *restful.Response) {
	req := &v1.CreatePolicyRequest{}

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.policy.CreatePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListPolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewListPolicyRequestFromHTTP(r)
	set, err := h.policy.ListPolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *permissionHandler) CheckPermission(r *restful.Request, w *restful.Response) {
	req := policy.NewCheckPermissionRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	perm, err := h.policy.CheckPermission(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, perm)
}
