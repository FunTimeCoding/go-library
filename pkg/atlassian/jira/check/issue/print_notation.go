package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(c *jira.Client) {
	r := report.New()

	for _, i := range query.Open(c) {
		r.AddItem(
			i.MonitorIdentifier,
			constant.WarningType,
			i.Summary,
			i.Link,
			i.Create,
		)
	}

	r.Print()
}
