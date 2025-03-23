package alert

import (
	"fmt"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
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

	c := alertmanager.NewEnvironment()

	if p.Notation {
		printNotation(c)

		return
	}

	if p.Rules {
		printRules(c, p.Firing, p.Old)

		return
	}

	f := statusOption.Color.Copy().Tag(
		tag.Link,
		tag.Runbook,
		tag.Category,
		tag.Emoji,
	)

	if p.Extended {
		f.Extended()
	}

	d := advanced_option.New()
	d.All = p.All
	d.CriticalOnly = p.Critical
	d.WarningOnly = p.Warning
	d.Suppressed = p.Suppressed
	d.Old = p.Old
	alerts, statistic := c.AlertsAdvanced(
		d,
		alert_enricher.New().Add(
			constant.KubernetesCronJobFailed,
			constant.Job,
			constant.Fail,
		),
		field_changer.New(),
		name_filter.New(true),
		label_filter.New(true),
	)

	for _, a := range alerts {
		fmt.Println(a.Format(f))

		if false {
			// TODO: More details: Annotations, rule
			fmt.Printf("  Raw: %+v\n", a.Raw)

			if r := c.Rule(a.Name); r != nil {
				if r.RawAlert != nil {
					fmt.Printf("  RawAlert: %+v\n", r.RawAlert)
				} else {
					fmt.Printf("  Rule: %+v\n", r)
				}
			}
		}
	}

	if !p.All && statistic.Relevant == 0 {
		fmt.Printf(
			"No relevant alerts, %d in total\n",
			statistic.Total,
		)
	}
}
