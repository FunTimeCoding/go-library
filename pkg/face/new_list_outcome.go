package face

func NewListOutcome(results []SearchResult) *ListOutcome {
	return &ListOutcome{Results: results}
}
