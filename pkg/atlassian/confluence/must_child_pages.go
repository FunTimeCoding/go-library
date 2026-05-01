package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustChildPages(
	space string,
	name string,
) []*page.Page {
	result, e := c.ChildPages(space, name)
	errors.PanicOnError(e)

	return result
}
