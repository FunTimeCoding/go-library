package silence

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
)

func printNotation(
	c *alertmanager.Client,
	o *option.Silence,
) {
	r := report.New()

	for _, e := range report.Trim(
		c.Silences(false),
		r,
		o.All,
		item.GoSilence,
	) {
		r.AddItem(
			item.GoSilence,
			e.MonitorIdentifier,
			constant.Warning,
			e.Rule,
			e.Link,
			e.Start,
		)
	}

	r.Print()
}
