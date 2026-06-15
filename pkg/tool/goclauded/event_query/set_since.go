package event_query

import "time"

func (q *Query) SetSince(v time.Time) *Query {
	q.since = v

	return q
}
