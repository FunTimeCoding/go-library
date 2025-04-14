package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"time"
)

func printNotation(
	c *alertmanager.Client,
	o *option.Silence,
) {
	r := report.New()
	silences := c.Silences(false)

	if !o.All && len(silences) > constant.NotationReport {
		silences = silences[0:constant.NotationReport]
		r.AddItem(
			fmt.Sprintf("%s-0", constant.SilencePrefix),
			constant.WarningType,
			fmt.Sprintf(
				"Too many silences, showing only the newest %d",
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	for _, s := range silences {
		r.AddItem(
			s.MonitorIdentifier,
			constant.WarningType,
			s.Rule,
			s.Link,
			s.Start,
		)
	}

	r.Print()
}
