package confluence

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) SpacesTreminioV2() []*models.SpaceSchemeV2 {
	result, r, e := c.treminioV2.Space.Bulk(
		c.context,
		&models.GetSpacesOptionSchemeV2{},
		"",
		25,
	)
	panicOnError(r, e)

	return result.Results
}
