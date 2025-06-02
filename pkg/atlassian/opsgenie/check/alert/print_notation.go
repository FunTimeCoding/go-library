package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	c *opsgenie.Client,
	o *option.Alert,
) {
	r := report.New()
	alerts := c.Open()

	if !o.All && len(alerts) > constant.NotationReport {
		alerts = alerts[0:constant.NotationReport]
		r.AddItem(
			constant.OpsgeniePrefix+"-0",
			constant.WarningLevel,
			fmt.Sprintf(
				"Too many Opsgenie alerts, showing only the newest %d",
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	for _, a := range alerts {
		var itemType string

		if a.Acknowledged {
			itemType = constant.WarningLevel
		} else {
			itemType = constant.ErrorLevel
		}

		r.AddItem(a.MonitorIdentifier, itemType, a.Name, a.Link, &a.Create)
	}

	r.Print()
}
