package report

import "github.com/funtimecoding/go-library/pkg/monitor/item"

func (r *Report) Add(i *item.Item) {
	r.Items = append(r.Items, i)
}
