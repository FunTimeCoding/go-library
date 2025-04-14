package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	alertmanagerConstant "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
	"time"
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

	if !o.All && len(relevant) > constant.NotationReport {
		relevant = relevant[0:constant.NotationReport]
		r.AddItem(
			fmt.Sprintf("%s-0", constant.AlertPrefix),
			constant.WarningType,
			fmt.Sprintf(
				"Too many Prometheus alerts, showing only the top %d",
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	for _, a := range relevant {
		var alertType string

		if a.Severity == alertmanagerConstant.CriticalSeverity {
			alertType = constant.ErrorType
		} else {
			alertType = constant.WarningType
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
