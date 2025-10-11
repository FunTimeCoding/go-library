package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SpacesTreminio() []*models.SpaceScheme {
	r, _, e := c.treminio.Space.Gets(
		c.context,
		&models.GetSpacesOptionScheme{},
		0,
		25,
	)
	errors.PanicOnError(e)

	return r.Results
}
