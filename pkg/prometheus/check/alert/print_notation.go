package alert

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	alertmanagerConstant "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
)

func printNotation(
	c *alertmanager.Client,
	o *option.Alert,
) {
	r := report.New()
	alerts, _ := c.Alerts(advanced_option.New())
	var relevant []*alert.Alert

	for _, a := range alerts {
		if !o.All && a.Severity == alertmanagerConstant.InfoSeverity {
			continue
		}

		relevant = append(relevant, a)
	}

	for _, a := range report.Trim(
		relevant,
		r,
		o.All,
		"Prometheus alerts",
		constant.AlertPrefix,
	) {
		var alertType string

		if a.Severity == alertmanagerConstant.CriticalSeverity {
			alertType = constant.ErrorLevel
		} else {
			alertType = constant.WarningLevel
		}

		r.AddItem(
			a.MonitorIdentifier,
			alertType,
			a.Name,
			a.Link,
			a.Start,
		)
	}

	r.Print()
}
