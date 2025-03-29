package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(c *opsgenie.Client) {
	r := report.New()

	for _, a := range c.Open() {
		var itemType string

		if a.Acknowledged {
			itemType = constant.WarningType
		} else {
			itemType = constant.ErrorType
		}

		r.AddItem(a.MonitorIdentifier, itemType, a.Name, a.Link)
	}

	r.Print()
}
