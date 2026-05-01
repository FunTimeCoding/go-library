package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPagesBySpace(identifier string) []*page.Page {
	result, e := c.PagesBySpace(identifier)
	errors.PanicOnError(e)

	return result
}
