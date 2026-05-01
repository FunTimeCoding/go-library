package confluence

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustAddComment(
	pageIdentifier string,
	body string,
) {
	errors.PanicOnError(c.AddComment(pageIdentifier, body))
}
