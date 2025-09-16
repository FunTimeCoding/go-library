package issue_enricher

func WithConcernFunction(v Slice) Option {
	return func(e *Enricher) {
		e.concerns = v
	}
}
