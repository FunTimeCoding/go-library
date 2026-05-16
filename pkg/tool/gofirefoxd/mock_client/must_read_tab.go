package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/firefox/content"
)

func (c *Client) MustReadTab(
	identifier int,
	raw bool,
) *content.Content {
	result, e := c.ReadTab(identifier, raw)
	errors.PanicOnError(e)

	return result
}
