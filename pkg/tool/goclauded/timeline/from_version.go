package timeline

import "fmt"

func FromVersion(
	name string,
	description string,
	changeType string,
	source string,
	changedAt string,
) *Entry {
	return &Entry{
		Timestamp: changedAt,
		Kind:      fmt.Sprintf("memory_%s", changeType),
		Actor:     source,
		Subject:   name,
		Detail:    description,
	}
}
