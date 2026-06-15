package event_query

import "time"

func (q *Query) SetBefore(v time.Time) *Query {
	q.before = v

	return q
}
