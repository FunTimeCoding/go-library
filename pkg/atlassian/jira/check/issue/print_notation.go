package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	jiraConstant "github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"time"
)

func printNotation(
	c *jira.Client,
	o *option.Issue,
) {
	r := report.New()
	issues := query.Open(
		c,
		environment.Get(jiraConstant.ProjectEnvironment, 1),
	)

	if !o.All && len(issues) > constant.NotationReport {
		issues = issues[0:constant.NotationReport]
		r.AddItem(
			constant.JiraPrefix+"-0",
			constant.WarningLevel,
			fmt.Sprintf(
				"Too many issues, showing only the newest %d",
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	for _, i := range issues {
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
