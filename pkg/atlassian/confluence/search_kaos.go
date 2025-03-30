package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SearchKaos(query string) []*content.Content {
	result, e := c.kaos.SearchContent(kaos.ContentSearchParameters{CQL: query})
	errors.PanicOnError(e)

	return content.NewSlice(result.Results, c.host)
}
