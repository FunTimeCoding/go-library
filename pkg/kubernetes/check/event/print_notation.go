package event

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"slices"
)

func printNotation(
	c *client.Client,
	o *option.Event,
) {
	r := report.New()
	var relevant []*event.Event

	for _, e := range c.EventsSimple(false, true) {
		if !o.All && slices.Contains(
			kubernetesConstant.IrrelevantEventReason,
			e.Reason,
		) {
			continue
		}

		relevant = append(relevant, e)
	}

	for _, e := range report.Trim(
		relevant,
		r,
		o.All,
		"Kubernetes events",
		constant.KubernetesEventPrefix,
	) {
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
