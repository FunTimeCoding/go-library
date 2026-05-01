package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustLabeled() []*page.Page {
	result, e := c.Labeled()
	errors.PanicOnError(e)

	return result
}
