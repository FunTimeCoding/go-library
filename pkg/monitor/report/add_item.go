package report

import "github.com/funtimecoding/go-library/pkg/monitor/item"

func (r *Report) AddItem(
	identifier string,
	itemType string,
	detail string,
	link string,
) {
	r.Add(item.New(identifier, itemType, detail, link))
}
