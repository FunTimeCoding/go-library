package event

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"slices"
	"time"
)

func printNotation(
	c *client.Client,
	o *option.Event,
) {
	r := report.New()
	events := c.EventsSimple(false, true)

	if !o.All && len(events) > constant.NotationReport {
		events = events[0:constant.NotationReport]
		r.AddItem(
			fmt.Sprintf("%s-0", constant.KubernetesEventPrefix),
			constant.WarningLevel,
			fmt.Sprintf(
				"Too many events, showing only the newest %d",
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	for _, e := range events {
		if !o.All && slices.Contains(
			kubernetesConstant.IrrelevantEventReason,
			e.Reason,
		) {
			continue
		}

		r.AddItem(
			e.MonitorIdentifier,
			constant.WarningLevel,
			e.Reason,
			"",
			e.Create,
		)
	}

	r.Print()
}
