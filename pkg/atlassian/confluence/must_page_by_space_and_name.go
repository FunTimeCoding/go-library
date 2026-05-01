package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustPageBySpaceAndName(
	spaceName string,
	name string,
) *page.Page {
	result, e := c.PageBySpaceAndName(spaceName, name)
	errors.PanicOnError(e)

	return result
}
