package alert

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func printNotation() {
	c := alertmanager.NewEnvironment()
	r := report.New()

	for _, a := range c.Alerts() {
		// TODO: If alert is critical, use ErrorType, otherwise use WarningType
		r.AddItem(
			a.MonitorIdentifier,
			constant.ErrorType,
			a.Name,
			a.Link,
		)
	}

	r.Print()
}
