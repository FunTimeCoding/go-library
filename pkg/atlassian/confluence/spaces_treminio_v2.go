package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SpacesTreminioV2() []*models.SpaceSchemeV2 {
	result, _, e := c.treminioV2.Space.Bulk(
		c.context,
		nil,
		"",
		25,
	)
	errors.PanicOnError(e)

	return result.Results
}
