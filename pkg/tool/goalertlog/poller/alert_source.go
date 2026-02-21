package poller

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

type AlertSource interface {
	Alerts(
		p *advanced_option.Alert,
	) ([]*alert.Alert, *statistic.Statistic)
}
