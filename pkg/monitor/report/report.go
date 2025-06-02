package report

import "github.com/funtimecoding/go-library/pkg/monitor/item"

type Report struct {
	Count int
	Items []*item.Item
	Note  []string
}
