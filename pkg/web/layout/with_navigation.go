package layout

import "maragu.dev/gomponents"

func (p *Page) WithNavigation(nodes ...gomponents.Node) *Page {
	p.navigation = nodes

	return p
}
