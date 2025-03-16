package silence

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func printNotation() {
	c := alertmanager.NewEnvironment()
	r := report.New()

	for _, s := range c.Silences(false) {
		r.AddItem(
			s.MonitorIdentifier,
			constant.WarningType,
			s.Rule,
			s.Link,
		)
	}

	r.Print()
}
