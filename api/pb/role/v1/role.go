package role

import (
	"ekube/api/pb/endpoint/v1"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
)

func (r *RoleSet) Add(item *Role) {
	r.Total++
	r.Items = append(r.Items, item)
}

func (r *Role) HasPermission(ep *endpoint.Endpoint) (*PermissionSpec, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		perm := r.Permissions[i]

		rok = perm.MatchResource(ep.ServiceId, ep.Entry.Resource)
		lok = perm.MatchLabel(ep.Entry.Labels)
		zap.L().Debugf("resource match: service_id: %r[target: %r] resource: %r[target: %r], result: %v",
			ep.ServiceId, perm.ServiceId, ep.Entry.Resource, perm.ResourceName, rok)
		zap.L().Debugf("label match: %v from [key: %v, value: %v, result: %v",
			ep.Entry.Labels, perm.LabelKey, perm.LabelValues, lok)
		if rok && lok {
			return perm, true, nil
		}
	}
	return nil, false, nil
}

func (r *Role) FullName() string {
	return fmt.Sprintf("%s@%s", r.Spec.Name, r.Spec.Scope)
}
