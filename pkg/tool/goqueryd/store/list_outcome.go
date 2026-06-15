package store

type ListOutcome struct {
	Results []SearchResult `json:"results"`
	Facets  []Facet        `json:"facets,omitempty"`
}
