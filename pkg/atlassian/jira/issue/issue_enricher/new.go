package issue_enricher

func New(
	concerns SliceResult,
	score FloatResult,
	commentNameFilter []string,
) *Enricher {
	return &Enricher{
		concerns:          concerns,
		score:             score,
		commentNameFilter: commentNameFilter,
	}
}
