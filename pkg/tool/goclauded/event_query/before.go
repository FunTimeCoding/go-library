package event_query

import "time"

func (q *Query) Before() time.Time { return q.before }
