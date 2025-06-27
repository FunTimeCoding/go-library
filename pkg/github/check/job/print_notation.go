package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(c *github.Client) {
	r := report.New()

	for _, e := range c.FailedRuns(false) {
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
