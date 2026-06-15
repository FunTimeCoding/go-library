package store

type SearchOutcome struct {
	Results  []SearchResult `json:"results"`
	Facets   []Facet        `json:"facets,omitempty"`
	Degraded bool           `json:"degraded,omitempty"`
	Cause    error          `json:"-"`
}
