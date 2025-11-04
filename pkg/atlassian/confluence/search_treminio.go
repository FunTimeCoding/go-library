package confluence

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) SearchTreminio(query string) []*models.SearchResultScheme {
	result, r, e := c.treminio.Search.Content(
		c.context,
		query,
		&models.SearchContentOptions{},
	)
	panicOnError(r, e)

	return result.Results
}
