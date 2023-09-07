package api

import (
	"ekube/internal/user"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
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
