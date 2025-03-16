package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
)

func (c *Client) Alerts() []*alert.Alert {
	response, e := c.client.Alert.GetAlerts(nil)
	errors.PanicOnError(e)
	result := alert.NewSlice(response.GetPayload(), c.host)

	for _, a := range result {
		a.Documentation = c.Documentation(a.Name)
	}

	return result
}
