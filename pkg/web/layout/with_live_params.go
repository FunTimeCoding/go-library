package layout

func (p *Page) WithLiveParams(query string) *Page {
	p.liveParams = query

	return p
}
