package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Status() *models.AlertmanagerStatus {
	response, e := c.client.General.GetStatus(nil)
	errors.PanicOnError(e)

	return response.GetPayload()
}
