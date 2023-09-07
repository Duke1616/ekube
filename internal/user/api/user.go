package api

import (
	"ekube/internal/user"
	"ekube/utils/response"
	"github.com/emicklei/go-restful/v3"
)

func (h *handler) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateUserRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	//tk := token.GetTokenFromRequest(r)
	//req.CreateBy = tk.UserId

	set, err := h.user.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ListUser(r *restful.Request, w *restful.Response) {
	req := user.NewListUserRequestFromHTTP(r.Request)
	set, err := h.user.ListUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
