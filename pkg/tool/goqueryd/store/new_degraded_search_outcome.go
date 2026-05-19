package store

func NewDegradedSearchOutcome(cause error) *SearchOutcome {
	return &SearchOutcome{Degraded: true, Cause: cause}
}
