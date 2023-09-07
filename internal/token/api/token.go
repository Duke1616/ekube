package api

import (
	"ekube/internal/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充用户的登录时的位置信息
	tk, err := h.token.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk.SetCookie(w)
	response.Success(w, tk)
}

func (h *handler) ValidateToken(r *restful.Request, w *restful.Response) {
	tk := token.GetAccessTokenFromHTTP(r.Request)
	req := token.NewValidateTokenRequest(tk)

	resp, err := h.token.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}
