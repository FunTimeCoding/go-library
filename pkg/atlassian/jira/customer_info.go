package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) CustomerInfo() *models.InfoScheme {
	result, _, e := c.service.Info.Get(c.context)
	errors.PanicOnError(e)

	return result
}
