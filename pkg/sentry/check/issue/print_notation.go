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

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		Plural,
		constant.SentryPrefix,
	) {
		r.AddItem(
			e.MonitorIdentifier,
			constant.ErrorLevel,
			e.Title,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
