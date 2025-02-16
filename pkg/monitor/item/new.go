package item

func New(
	identifier string,
	itemType string,
	detail string,
	link string,
) *Item {
	return &Item{
		Identifier: identifier,
		Type:       itemType,
		Detail:     detail,
		Link:       link,
	}
}
