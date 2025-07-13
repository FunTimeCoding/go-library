package issue_enricher

func New(o ...OptionFunc) *Enricher {
	result := &Enricher{}

	for _, p := range o {
		p(result)
	}

	return result
}
