package store

func NewSearchOutcome(results []SearchResult) *SearchOutcome {
	return &SearchOutcome{Results: results}
}
