package job

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*job.Job,
	o *option.Job,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		Plural,
		monitor.GitLabPrefix,
	) {
		var level string

		if e.Fail() {
			level = monitor.ErrorLevel
		} else {
			level = monitor.InformationLevel
		}

		r.AddItem(
			e.MonitorIdentifier,
			level,
			e.Format(f),
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
