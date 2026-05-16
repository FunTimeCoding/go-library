package checklist_item

func New(
	index int,
	text string,
	checked bool,
) *Item {
	return &Item{Index: index, Text: text, Checked: checked}
}
