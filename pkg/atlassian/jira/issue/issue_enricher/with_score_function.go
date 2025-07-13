package issue_enricher

func WithScoreFunction(v FloatResult) OptionFunc {
	return func(e *Enricher) {
		e.score = v
	}
}
