package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) RequestTypes(
	desk int,
	group int,
) *models.ProjectRequestTypePageScheme {
	result, _, e := c.service.Request.Type.Gets(
		c.context,
		desk,
		group,
		0,
		50,
	)
	errors.PanicOnError(e)

	return result
}
