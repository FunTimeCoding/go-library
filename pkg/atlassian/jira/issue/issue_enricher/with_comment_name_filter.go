package issue_enricher

func WithCommentNameFilter(v []string) OptionFunc {
	return func(e *Enricher) {
		e.commentNameFilter = v
	}
}
