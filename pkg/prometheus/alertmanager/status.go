package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Status() *models.AlertmanagerStatus {
	r, e := c.client.General.GetStatus(&general.GetStatusParams{})
	errors.PanicOnError(e)

	return r.GetPayload()
}
