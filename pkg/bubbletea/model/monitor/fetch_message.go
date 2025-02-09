package monitor

import "github.com/funtimecoding/go-library/pkg/monitor/item"

type FetchMessage struct {
	Items []*item.Item
}
