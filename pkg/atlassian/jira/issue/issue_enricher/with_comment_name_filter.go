package issue_enricher

func WithCommentNameFilter(v []string) Option {
	return func(e *Enricher) {
		e.commentNameFilter = v
	}
}
