package alert

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
)

func printNotation(c *alertmanager.Client) {
	r := report.New()
	alerts, _ := c.Alerts(advanced_option.New())

	for _, a := range alerts {
		// TODO: If alert is critical, use ErrorType, otherwise use WarningType
		r.AddItem(
			a.MonitorIdentifier,
			constant.ErrorType,
			a.Name,
			a.Link,
			a.Start,
		)
	}

	r.Print()
}
