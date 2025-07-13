package version

import (
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/go-library/pkg/go_mod/project"
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*project.Project,
	o *option.Version,
) {
	r := report.New()
	f := go_mod.Format

	for _, e := range report.Trim(v, r, o.All, Plural, monitor.GitPrefix) {
		var level string

		if e.HasConcerns() {
			level = monitor.WarningLevel
		} else {
			level = monitor.InformationLevel
		}

		r.AddItem(
			e.MonitorIdentifier,
			level,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
