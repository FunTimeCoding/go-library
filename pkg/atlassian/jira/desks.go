package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Desks() *models.ServiceDeskPageScheme {
	result, _, e := c.service.ServiceDesk.Gets(c.context, 0, 50)
	errors.PanicOnError(e)

	return result
}
