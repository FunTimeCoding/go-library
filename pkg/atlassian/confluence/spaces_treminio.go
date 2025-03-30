package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SpacesTreminio() []*models.SpaceScheme {
	result, _, e := c.treminio.Space.Gets(
		c.context,
		nil,
		0,
		25,
	)
	errors.PanicOnError(e)

	return result.Results
}
