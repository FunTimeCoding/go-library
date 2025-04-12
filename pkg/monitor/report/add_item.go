package report

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"time"
)

func (r *Report) AddItem(
	identifier string,
	itemType string,
	detail string,
	link string,
	t *time.Time,
) {
	r.Add(item.New(identifier, itemType, detail, link, t))
}
