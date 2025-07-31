package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGenie,
	) {
		var s constant.Severity

		if e.Acknowledged {
			s = constant.Warning
		} else {
			s = constant.Critical
		}

		r.AddItem(
			item.GoGenie,
			e.MonitorIdentifier,
			s,
			e.Name,
			e.Link,
			&e.Create,
		)
	}

	r.Print()
}
