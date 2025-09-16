package issue_enricher

func WithScoreFunction(v Float) Option {
	return func(e *Enricher) {
		e.score = v
	}
}
