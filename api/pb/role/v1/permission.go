package role

import (
	"fmt"
	"hash/fnv"
)

func (s *PermissionSet) PermissionSpecs() (items []*PermissionSpec) {
	for i := range s.Items {
		item := s.Items[i]
		items = append(items, item.Spec)
	}
	return
}

func (req *AddPermissionToRoleRequest) Length() int {
	return len(req.Permissions)
}

func (p *PermissionSpec) HashID(roleId string) string {
	h := fnv.New32a()

	h.Write([]byte(roleId + p.Effect.String() + p.ServiceId + p.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}

func (p *PermissionSpec) MatchResource(serviceID, resourceName string) bool {
	// 服务匹配
	if p.ServiceId != "*" && p.ServiceId != serviceID {
		return false
	}

	// 资源匹配
	if p.ResourceName != "*" && p.ResourceName != resourceName {
		return false
	}

	return true
}

// MatchLabel 匹配Label
func (p *PermissionSpec) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		// 匹配key
		if p.LabelKey == "*" || p.LabelKey == k {
			// 匹配value
			if p.isMatchAllValue() {
				return true
			}
			for i := range p.LabelValues {
				if p.LabelValues[i] == v {
					return true
				}
			}
		}
	}

	return false
}

func (p *PermissionSpec) isMatchAllValue() bool {
	if p.MatchAll {
		return true
	}

	for i := range p.LabelValues {
		if p.LabelValues[i] == "*" {
			return true
		}
	}

	return false
}
