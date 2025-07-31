package job

import (
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(
	v []*run.Run,
	o *option.Job,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGitHub,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Critical
		} else {
			s = constant.Information
		}

		r.AddItem(
			item.GoGitHub,
			e.MonitorIdentifier,
			s,
			e.Format(Notation),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
