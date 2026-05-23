package usage

import "time"

type block struct {
	start   time.Time
	end     time.Time
	entries []*timestamped
}
