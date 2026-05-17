package layout

func (p *Page) WithPath(path string) *Page {
	p.path = path

	return p
}
