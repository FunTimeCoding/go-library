package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (c *Client) Alerts(
	_ *advanced_option.Alert,
) ([]*alert.Alert, *statistic.Statistic) {
	return c.alerts, statistic.New()
}
