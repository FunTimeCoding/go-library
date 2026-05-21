package ensure_result

import "time"

func NewExisting(
	name string,
	lastSeen time.Time,
) *Result {
	return &Result{Callsign: name, LastSeen: lastSeen}
}
