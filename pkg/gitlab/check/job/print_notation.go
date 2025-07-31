package job

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*job.Job,
	o *option.Job,
) {
	r := report.New()
	f := constant.CheckFormat

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGitLab,
	) {
		var s monitor.Severity

		if e.Fail() {
			s = monitor.Critical
		} else {
			s = monitor.Information
		}

		r.AddItem(
			item.GoGitLab,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
