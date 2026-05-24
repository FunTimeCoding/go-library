package layout

func (p *Page) Clone() *Page {
	return &Page{
		identity:   p.identity,
		theme:      p.theme,
		style:      p.style,
		brandNode:  p.brandNode,
		items:      p.items,
		navigation: p.navigation,
		script:       p.script,
		footer:       p.footer,
		liveEndpoint: p.liveEndpoint,
	}
}
