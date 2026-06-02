package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert_filter"
)

func (c *Client) MustAlerts(
	_ *advanced_option.Alert,
	_ *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic) {
	return c.alerts, statistic.New()
}
