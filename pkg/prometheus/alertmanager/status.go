package alertmanager

import (
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Status() (*models.AlertmanagerStatus, error) {
	r, e := c.client.General.GetStatus(&general.GetStatusParams{})

	if e != nil {
		return nil, e
	}

	return r.GetPayload(), nil
}
