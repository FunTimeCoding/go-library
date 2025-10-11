package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/client/receiver"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Receivers() []*models.Receiver {
	r, e := c.client.Receiver.GetReceivers(
		&receiver.GetReceiversParams{},
	)
	errors.PanicOnError(e)

	return r.GetPayload()
}
