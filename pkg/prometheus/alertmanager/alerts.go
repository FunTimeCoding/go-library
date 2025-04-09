package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (c *Client) Alerts(p *advanced_option.Alert) ([]*alert.Alert, *statistic.Statistic) {
	response, e := c.client.Alert.GetAlerts(nil)
	errors.PanicOnError(e)

	return c.processor(p).Process(alert.NewSlice(response.GetPayload(), c.host))
}
