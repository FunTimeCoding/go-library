package ensure_result

import "time"

type Result struct {
	Callsign string
	New      bool
	LastSeen time.Time
}
