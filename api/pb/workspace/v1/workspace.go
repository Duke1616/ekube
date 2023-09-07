package workspace

func (w *WorkspaceSet) Add(item *Workspace) {
	w.Items = append(w.Items, item)
}
