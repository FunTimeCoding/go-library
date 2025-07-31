package issue

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/sentry/issue"
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
		item.GoSentry,
	) {
		r.AddItem(
			item.GoSentry,
			e.MonitorIdentifier,
			constant.Critical,
			e.Title,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
