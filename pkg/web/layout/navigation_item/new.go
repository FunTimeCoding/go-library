package navigation_item

func New(
	path string,
	label string,
) *Item {
	return &Item{path: path, label: label}
}
