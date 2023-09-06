package tools

import (
	"ekube/third_party/ekube/pb/page"
	"net/http"
	"strconv"
)

const (
	// DefaultPageSize 默认分页大小
	DefaultPageSize = 10
	// DefaultPageNumber 默认页号
	DefaultPageNumber = 1
)

func NewDefaultPageRequest() *page.PageRequest {
	return NewPageRequest(DefaultPageSize, DefaultPageNumber)
}

// NewPageRequest 实例化
func NewPageRequest(ps uint, pn uint) *page.PageRequest {
	return &page.PageRequest{
		PageSize:   uint64(ps),
		PageNumber: uint64(pn),
	}
}

func NewPageRequestFromHTTP(req *http.Request) *page.PageRequest {
	qs := req.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	os := qs.Get("offset")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)
	osInt64, _ := strconv.ParseInt(os, 10, 64)

	if psUint64 == 0 {
		psUint64 = DefaultPageSize
	}
	if pnUint64 == 0 {
		pnUint64 = DefaultPageNumber
	}

	return &page.PageRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Offset:     osInt64,
	}
}
