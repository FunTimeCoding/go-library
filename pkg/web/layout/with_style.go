package layout

func (p *Page) WithStyle(style string) *Page {
	p.style = style

	return p
}
