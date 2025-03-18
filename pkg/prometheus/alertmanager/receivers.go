package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Receivers() []*models.Receiver {
	result, e := c.client.Receiver.GetReceivers(nil, nil)
	errors.PanicOnError(e)

	return result.GetPayload()
}
