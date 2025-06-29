package event

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"slices"
)

func printNotation(
	c *client.Client,
	o *option.Event,
) {
	r := report.New()

	for _, e := range report.Trim(
		c.EventsSimple(false, true),
		r,
		o.All,
		"Kubernetes events",
		constant.KubernetesEventPrefix,
	) {
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
