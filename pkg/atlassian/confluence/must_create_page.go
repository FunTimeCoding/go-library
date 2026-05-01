package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustCreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) *page.Page {
	result, e := c.CreatePage(spaceIdentifier, parentIdentifier, title, markdown)
	errors.PanicOnError(e)

	return result
}
