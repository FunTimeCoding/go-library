package navigation_item

func NewExternal(
	path string,
	label string,
) *Item {
	return &Item{path: path, label: label, external: true}
}
