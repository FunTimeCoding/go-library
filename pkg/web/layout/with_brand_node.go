package layout

import "maragu.dev/gomponents"

func (p *Page) WithBrandNode(node gomponents.Node) *Page {
	p.brandNode = node

	return p
}
