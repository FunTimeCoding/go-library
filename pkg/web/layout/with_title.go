package layout

func (p *Page) WithTitle(title string) *Page {
	p.title = title

	return p
}
