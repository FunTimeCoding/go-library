package key_reader

import "time"

type state struct {
	press   time.Time
	lastKey time.Time
	holding bool
}
