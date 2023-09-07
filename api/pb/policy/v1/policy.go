package policy

import (
	"fmt"
	"hash/fnv"
)

// RoleNames 获取用户的角色
func (p *PolicySet) RoleNames() (rns []string) {
	for i := range p.Items {
		item := p.Items[i]
		rns = append(rns, item.Role.FullName())
	}
	return
}

func (p *Policy) GenID() {
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		p.Spec.Workspace, p.Spec.Namespace, p.Spec.UserId, p.Spec.RoleId)

	h.Write([]byte(hashedStr))
	p.Meta.Id = fmt.Sprintf("%x", h.Sum32())
}

func (p *PolicySet) Add(item *Policy) {
	p.Items = append(p.Items, item)
}
