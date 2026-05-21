package resolve_result

func New(matches []*Match) *Result {
	return &Result{Matches: matches}
}
