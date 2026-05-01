package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPage(identifier string) *page.Page {
	result, e := c.Page(identifier)
	errors.PanicOnError(e)

	return result
}
