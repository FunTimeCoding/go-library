package layout

func (p *Page) WithScript(sources ...string) *Page {
	p.script = sources

	return p
}
