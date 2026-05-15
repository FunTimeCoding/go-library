package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) MustReadScreen(identifier string) session.Screen {
	result, e := c.ReadScreen(identifier)
	errors.PanicOnError(e)

	return result
}
