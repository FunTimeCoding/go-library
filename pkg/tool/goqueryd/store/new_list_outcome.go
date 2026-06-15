package store

func NewListOutcome(
	results []SearchResult,
	facets []Facet,
) *ListOutcome {
	return &ListOutcome{Results: results, Facets: facets}
}
