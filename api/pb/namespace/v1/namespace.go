package namespace

func (n *NamespaceSet) Add(item *Namespace) {
	n.Items = append(n.Items, item)
}
