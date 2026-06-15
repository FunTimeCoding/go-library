package event_query

func (q *Query) SetOffset(v int) *Query {
	q.offset = v

	return q
}
