package confluence

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) SpacesTreminio() []*models.SpaceScheme {
	result, r, e := c.treminio.Space.Gets(
		c.context,
		&models.GetSpacesOptionScheme{},
		0,
		25,
	)
	panicOnError(r, e)

	return result.Results
}
