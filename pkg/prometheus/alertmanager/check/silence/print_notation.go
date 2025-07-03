package silence

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
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
		Plural,
		constant.SilencePrefix,
	) {
		r.AddItem(
			e.MonitorIdentifier,
			constant.WarningLevel,
			e.Rule,
			e.Link,
			e.Start,
		)
	}

	r.Print()
}
