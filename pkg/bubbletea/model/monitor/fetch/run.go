package fetch

import (
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Run(c string) []*item.Item {
	if run.CommandExists(c) {
		if m := monitor.Run(c); m != nil {
			return m.Items
		}
	}

	return nil
}
