package issue

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func printNotation(c *sentry.Client) {
	r := report.New()

	for _, i := range c.IssuesSimple() {
		r.AddItem(
			i.MonitorIdentifier,
			constant.ErrorType,
			i.Title,
			i.Link,
			i.Create,
		)
	}

	r.Print()
}
