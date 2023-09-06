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

func (req *PermissionSpec) HashID(roleId string) string {
	h := fnv.New32a()

	h.Write([]byte(roleId + req.Effect.String() + req.ServiceId + req.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}
