package user

func (r *UserSet) Add(item *User) {
	r.Items = append(r.Items, item)
}
