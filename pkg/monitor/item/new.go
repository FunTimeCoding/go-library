package item

import "time"

func New(
	identifier string,
	level string,
	detail string,
	link string,
	t *time.Time,
) *Item {
	return &Item{
		Identifier: identifier,
		Level:      level,
		Detail:     detail,
		Link:       link,
		Create:     t,
	}
}
