package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Groups() models.AlertGroups {
	response, e := c.client.Alertgroup.GetAlertGroups(nil, nil)
	errors.PanicOnError(e)

	return response.GetPayload()
}
