package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
)

func Check(o *option.Alert) {
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
	d := advanced_option.New()
	d.All = o.All
	d.CriticalOnly = o.Critical
	d.WarningOnly = o.Warning
	d.Suppressed = o.Suppressed
	d.Old = o.Old
	alerts, statistic := c.Alerts(d)

	if o.Notation {
		printNotation(alerts, o)

		return
	}

	if o.Rules {
		printRules(c, o.Firing, o.Old)

		return
	}

	f := constant.Format.Copy()

	if o.Extended {
		f.Extended()
	}

	m := c.Rules()

	for _, a := range alerts {
		// TODO: Rule details
		fmt.Println(a.Format(f))

		if r := m.Find(a.Name); r != nil {
			fmt.Printf("  Rule: %s\n", r.Format(f))
		}
	}

	if !o.All && statistic.Relevant == 0 {
		fmt.Printf(
			"No relevant %s, %d in total\n",
			item.GoAlert.Plural,
			statistic.Total,
		)
	}
}
