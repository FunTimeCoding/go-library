package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPagesByLabel(labelIdentifier string) []*page.Page {
	result, e := c.PagesByLabel(labelIdentifier)
	errors.PanicOnError(e)

	return result
}
