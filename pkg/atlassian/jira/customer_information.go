package jira

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) CustomerInformation() (*models.InfoScheme, error) {
	result, _, e := c.service.Info.Get(c.context)

	return result, e
}
