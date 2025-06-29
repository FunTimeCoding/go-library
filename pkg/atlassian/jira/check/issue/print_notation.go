package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func printNotation(
	c *jira.Client,
	o *option.Issue,
) {
	r := report.New()

	for _, i := range report.Trim(
		query.Open(
			c,
			environment.Get(constant.ProjectEnvironment),
		),
		r,
		o.All,
		"Jira issues",
		monitorConstant.JiraPrefix,
	) {
		r.AddItem(
			i.MonitorIdentifier,
			monitorConstant.WarningLevel,
			i.Summary,
			i.Link,
			i.Create,
		)
	}

	r.Print()
}
