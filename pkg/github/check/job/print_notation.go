package job

import (
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
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
		Plural,
		constant.GitHubPrefix,
	) {
		var level string

		if e.HasConcerns() {
			level = constant.ErrorLevel
		} else {
			level = constant.InformationLevel
		}

		r.AddItem(
			e.MonitorIdentifier,
			level,
			e.Format(Notation),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
