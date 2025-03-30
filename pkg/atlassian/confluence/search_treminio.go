package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SearchTreminio(query string) []*models.SearchResultScheme {
	result, _, e := c.treminio.Search.Content(c.context, query, nil)
	errors.PanicOnError(e)

	return result.Results
}
