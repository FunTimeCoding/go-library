package issue

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/sentry/issue"
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
		"Sentry issues",
		constant.SentryPrefix,
	) {
		r.AddItem(
			i.MonitorIdentifier,
			constant.ErrorLevel,
			i.Title,
			i.Link,
			i.Create,
		)
	}

	r.Print()
}
