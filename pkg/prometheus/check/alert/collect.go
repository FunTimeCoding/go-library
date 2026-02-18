package alert

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
)

func collect(
	c *alertmanager.Client,
	o *option.Alert,
) ([]*alert.Alert, *statistic.Statistic) {
	d := advanced_option.New()
	d.All = o.All
	d.CriticalOnly = o.Critical
	d.WarningOnly = o.Warning
	d.Suppressed = o.Suppressed
	d.Old = o.Old

	return c.Alerts(d)
}
