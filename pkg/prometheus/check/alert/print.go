package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
)

func Print(p *option.Alert) {
	// TODO: Use in gomonitor
	//  Highlight new alerts
	//  Press key to add 10 minute silence in gomonitor
	//  Press key to delete silence in gomonitor
	// TODO: How to notify reliably on the desktop?
	//  Rely on Prometheus mails? Noise of other mails would be a problem.
	//   For mail, a mail client that filters would help
	//   How to get alert notifications as events?
	//    If events not available, store current alerts in memory and only notify on new alerts

	c := internal.Alertmanager()

	if p.Notation {
		printNotation(c)

		return
	}

	if p.Rules {
		printRules(c, p.Firing, p.Old)

		return
	}

	f := constant.Format.Copy()

	if p.Extended {
		f.Extended()
	}

	d := advanced_option.New()
	d.All = p.All
	d.CriticalOnly = p.Critical
	d.WarningOnly = p.Warning
	d.Suppressed = p.Suppressed
	d.Old = p.Old
	alerts, statistic := c.Alerts(d)

	m := c.Rules()

	for _, a := range alerts {
		// TODO: Rule details
		fmt.Println(a.Format(f))

		if r := m.Find(a.Name); r != nil {
			fmt.Printf("  Rule: %s\n", r.Format(f))
		}
	}

	if !p.All && statistic.Relevant == 0 {
		fmt.Printf(
			"No relevant alerts, %d in total\n",
			statistic.Total,
		)
	}
}
