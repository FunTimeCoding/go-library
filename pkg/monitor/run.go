package monitor

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Run(name string) []*item.Item {
	r := run.New()
	r.Panic = false
	r.Start(name)
	var result []*item.Item
	notation.DecodeStrict(r.OutputString, &result, true)

	return result
}
