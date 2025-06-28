package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*job.Job,
	o *option.Job,
) {
	r := report.New()

	if !o.All && len(v) > monitorConstant.NotationReport {
		v = v[0:monitorConstant.NotationReport]
		r.AddItem(
			monitorConstant.GitLabPrefix+"-0",
			monitorConstant.WarningLevel,
			fmt.Sprintf(
				"Too many failed jobs, showing only the newest %d",
				monitorConstant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	f := constant.Format

	for _, e := range v {
		r.AddItem(
			e.MonitorIdentifier,
			monitorConstant.WarningLevel,
			e.Format(f),
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
