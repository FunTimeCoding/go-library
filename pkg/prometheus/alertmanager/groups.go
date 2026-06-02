package alertmanager

import "github.com/prometheus/alertmanager/api/v2/models"

func (c *Client) Groups() (models.AlertGroups, error) {
	response, e := c.client.Alertgroup.GetAlertGroups(nil, nil)

	if e != nil {
		return nil, e
	}

	return response.GetPayload(), nil
}
