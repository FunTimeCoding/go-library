package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*issue.Issue,
	o *option.Issue,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoJira,
	) {
		r.AddItem(
			item.GoJira,
			e.MonitorIdentifier,
			constant.Warning,
			e.Summary,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
