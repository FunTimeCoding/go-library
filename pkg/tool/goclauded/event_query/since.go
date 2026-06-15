package event_query

import "time"

func (q *Query) Since() time.Time { return q.since }
