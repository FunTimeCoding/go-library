package layout

func (p *Page) WithTheme(theme string) *Page {
	p.theme = theme

	return p
}
