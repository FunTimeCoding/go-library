package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPages() []*page.Page {
	result, e := c.Pages()
	errors.PanicOnError(e)

	return result
}
