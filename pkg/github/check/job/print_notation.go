package job

import (
	"fmt"
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
		"GitHub jobs",
		constant.GitHubPrefix,
	) {
		r.AddItem(
			e.MonitorIdentifier,
			constant.WarningLevel,
			fmt.Sprintf("%s fail", e.Repository().FullName),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
