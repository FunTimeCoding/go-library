package jira

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) Desks() (*models.ServiceDeskPageScheme, error) {
	result, _, e := c.service.ServiceDesk.Gets(c.context, 0, 50)

	return result, e
}
