package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*issue.Issue,
	o *option.Issue,
) {
	r := report.New()

	for _, i := range report.Trim(
		v,
		r,
		o.All,
		Plural,
		constant.JiraPrefix,
	) {
		r.AddItem(
			i.MonitorIdentifier,
			constant.WarningLevel,
			i.Summary,
			i.Link,
			i.Create,
		)
	}

	r.Print()
}
