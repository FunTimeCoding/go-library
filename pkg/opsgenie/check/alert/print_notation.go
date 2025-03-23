package alert

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/opsgenie"
)

func printNotation(c *opsgenie.Client) {
	r := report.New()

	for _, a := range c.Open() {
		// TODO: If alert is P1-P3, use ErrorType, otherwise use WarningType
		r.AddItem(
			a.MonitorIdentifier,
			constant.ErrorType,
			a.Name,
			a.Link,
		)
	}

	r.Print()
}
