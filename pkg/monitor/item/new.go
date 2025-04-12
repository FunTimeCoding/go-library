package item

import "time"

func New(
	identifier string,
	itemType string,
	detail string,
	link string,
	t *time.Time,
) *Item {
	return &Item{
		Identifier: identifier,
		Type:       itemType,
		Detail:     detail,
		Link:       link,
		Create:     t,
	}
}
