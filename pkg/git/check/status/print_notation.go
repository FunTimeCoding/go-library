package status

import (
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*repository.Repository,
	o *option.Status,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, Plural, constant.GitPrefix) {
		r.AddItem(
			e.MonitorIdentifier,
			constant.WarningLevel,
			e.Format(),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
