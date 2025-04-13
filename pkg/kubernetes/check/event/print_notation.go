package event

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"slices"
)

func printNotation(
	c *client.Client,
	o *option.Event,
) {
	r := report.New()

	for _, e := range c.EventsSimple(false, true) {
		if !o.All && slices.Contains(
			constant.IrrelevantEventReason,
			e.Reason,
		) {
			continue
		}

		r.AddItem(
			e.Name,
			monitorConstant.WarningType,
			e.Reason,
			"",
			e.Create,
		)
	}

	r.Print()
}
