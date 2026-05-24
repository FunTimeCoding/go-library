package layout

func (p *Page) WithLiveEndpoint(path string) *Page {
	p.liveEndpoint = path

	return p
}
