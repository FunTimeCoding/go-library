package event_query

func (q *Query) SetKinds(v []string) *Query {
	q.kinds = v

	return q
}
