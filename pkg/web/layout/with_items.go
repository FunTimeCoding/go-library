package layout

import "github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"

func (p *Page) WithItems(
	items ...*navigation_item.Item,
) *Page {
	p.items = items

	return p
}
