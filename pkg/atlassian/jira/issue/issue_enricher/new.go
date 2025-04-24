package issue_enricher

func New(
	concerns SliceResult,
	score FloatResult,
) *Enricher {
	return &Enricher{concerns: concerns, score: score}
}
