package api

import (
	"ekube/internal/endpoint"
	"ekube/utils/response"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
)

// RegistryEndpoint 注册Endpint
func (h *handler) RegistryEndpoint(r *restful.Request, w *restful.Response) {
	req := endpoint.NewDefaultRegistryRequest()
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	_, err := h.service.RegistryEndpoint(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
}

func (h *handler) ListEndpoints(r *restful.Request, w *restful.Response) {
	req := endpoint.NewQueryEndpointRequestFromHTTP(r.Request)

	set, err := h.service.ListEndpoints(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeEndpoint(r *restful.Request, w *restful.Response) {
	req := endpoint.NewDescribeEndpointRequestWithID(r.PathParameter("id"))

	d, err := h.service.DescribeEndpoint(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
