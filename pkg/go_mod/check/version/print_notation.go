package version

import (
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/go-library/pkg/go_mod/project"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*project.Project,
	o *option.Version,
) {
	r := report.New()
	f := go_mod.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGitStatus,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Warning
		} else {
			s = constant.Information
		}

		r.AddItem(
			item.GoGitStatus,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
