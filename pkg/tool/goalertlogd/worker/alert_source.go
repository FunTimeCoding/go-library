package worker

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert_filter"
)

type AlertSource interface {
	MustAlerts(
		p *advanced_option.Alert,
		f *alert_filter.Filter,
	) ([]*alert.Alert, *statistic.Statistic)
}
