package event_query

import "time"

type Query struct {
	since  time.Time
	before time.Time
	kinds  []string
	limit  int
	offset int
}
