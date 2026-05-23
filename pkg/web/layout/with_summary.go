package layout

func (p *Page) WithSummary(items ...string) *Page {
	p.summary = items

	return p
}
