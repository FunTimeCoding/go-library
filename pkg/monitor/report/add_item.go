package report

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"time"
)

func (r *Report) AddItem(
	identifier string,
	level string,
	detail string,
	link string,
	t *time.Time,
) {
	r.Add(item.New(identifier, level, detail, link, t))
}
