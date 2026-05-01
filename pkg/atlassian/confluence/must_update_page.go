package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustUpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) *page.Page {
	result, e := c.UpdatePage(identifier, title, markdown, message)
	errors.PanicOnError(e)

	return result
}
