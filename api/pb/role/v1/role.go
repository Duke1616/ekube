package role

func (s *RoleSet) Add(item *Role) {
	s.Total++
	s.Items = append(s.Items, item)
}
