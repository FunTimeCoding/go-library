package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPagesBySpaceName(n string) []*page.Page {
	result, e := c.PagesBySpaceName(n)
	errors.PanicOnError(e)

	return result
}
