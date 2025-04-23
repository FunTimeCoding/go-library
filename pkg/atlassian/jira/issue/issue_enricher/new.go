package issue_enricher

func New(v Validator) *Enricher {
	return &Enricher{validator: v}
}
