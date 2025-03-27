package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"
	"github.com/funtimecoding/go-library/pkg/errors"
	virtomize "github.com/virtomize/confluence-go-api"
)

func (c *Client) SearchVirtomize(query string) []*content.Content {
	result, e := c.virtomize.Search(virtomize.SearchQuery{CQL: query})
	errors.PanicOnError(e)

	return content.NewSliceVirtomize(result.Results)
}
