package issue_enricher

func WithConcernFunction(v SliceResult) OptionFunc {
	return func(e *Enricher) {
		e.concerns = v
	}
}
