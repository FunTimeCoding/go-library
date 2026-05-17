package layout

import "maragu.dev/gomponents"

func (p *Page) WithContent(nodes ...gomponents.Node) *Page {
	p.content = nodes

	return p
}
