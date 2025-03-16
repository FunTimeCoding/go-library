package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
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
	f := format.Color.Copy().Tag(tag.Link, tag.Emoji)
	var relevant int
	alerts := c.Alerts()

	for _, a := range alerts {
		if !p.All && a.Severity == constant.NoneSeverity {
			continue
		}

		relevant++
		fmt.Println(a.Format(f))
	}

	if !p.All && relevant == 0 {
		fmt.Printf("No relevant alerts, %d in total\n", len(alerts))
	}
}
