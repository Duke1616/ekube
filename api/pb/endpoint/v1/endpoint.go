package endpoint

import (
	"ekube/tools"
	"github.com/infraboard/mcube/types/ftime"
)

func (req *RegistryRequest) Endpoints() []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:        tools.GenHashID(req.ServiceId, req.Entries[i].Path),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: req.ServiceId,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}
