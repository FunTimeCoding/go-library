package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SpacesTreminioV2() []*models.SpaceSchemeV2 {
	r, _, e := c.treminioV2.Space.Bulk(
		c.context,
		&models.GetSpacesOptionSchemeV2{},
		"",
		25,
	)
	errors.PanicOnError(e)

	return r.Results
}
