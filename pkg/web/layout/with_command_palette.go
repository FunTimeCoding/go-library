package layout

func (p *Page) WithCommandPalette(endpoint string) *Page {
	p.paletteEndpoint = endpoint

	return p
}
