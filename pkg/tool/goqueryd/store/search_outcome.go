package store

type SearchOutcome struct {
	Results  []SearchResult
	Degraded bool
	Cause    error
}
