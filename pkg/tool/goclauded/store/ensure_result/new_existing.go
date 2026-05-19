package ensure_result

import "time"

func NewExisting(
	name string,
	lastSeen time.Time,
) *Result {
	return &Result{Name: name, LastSeen: lastSeen}
}
