package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/parameter"
)

func Print(p *parameter.Alert) {
	// TODO: Use in gomonitor
	//  Highlight new alerts
	//  Press key to add 10 minute silence in gomonitor
	//  Press key to delete silence in gomonitor
	// TODO: How to notify reliably on the desktop?
	//  Rely on Prometheus mails? Noise of other mails would be a problem.
	//   For mail, a mail client that filters would help
	//   How to get alert notifications as events?
	//    If events not available, store current alerts in memory and only notify on new alerts

	if p.Notation {
		printNotation()

		return
	}

	c := alertmanager.NewEnvironment()
	f := format.Color.Copy().Tag(tag.Link, tag.Documentation, tag.Emoji)

	if p.Extended {
		f.Extended()
	}

	p2 := advanced_parameter.New()
	p2.All = p.All
	p2.CriticalOnly = p.Critical
	p2.WarningOnly = p.Warning
	p2.Suppressed = p.Suppressed
	p2.Old = p.Old
	alerts, statistic := c.AlertsAdvanced(
		p2,
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
