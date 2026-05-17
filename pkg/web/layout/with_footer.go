package layout

import "maragu.dev/gomponents"

func (p *Page) WithFooter(nodes ...gomponents.Node) *Page {
	p.footer = nodes

	return p
}
