package status

import (
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*repository.Repository,
	o *option.Status,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(v, r, o.All, Plural, monitor.GitPrefix) {
		r.AddItem(
			e.MonitorIdentifier,
			monitor.WarningLevel,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
