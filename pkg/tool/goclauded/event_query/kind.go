package event_query

func (q *Query) Kind(v string) *Query {
	q.kinds = []string{v}

	return q
}
