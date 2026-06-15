package event_query

func (q *Query) SetLimit(v int) *Query {
	q.limit = v

	return q
}
