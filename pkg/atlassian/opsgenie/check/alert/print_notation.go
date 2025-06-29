package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()

	for _, a := range report.Trim(
		v,
		r,
		o.All,
		"Opsgenie alerts",
		constant.OpsgeniePrefix,
	) {
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
