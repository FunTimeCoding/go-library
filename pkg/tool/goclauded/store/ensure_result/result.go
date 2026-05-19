package ensure_result

import "time"

type Result struct {
	Name     string
	New      bool
	LastSeen time.Time
}
