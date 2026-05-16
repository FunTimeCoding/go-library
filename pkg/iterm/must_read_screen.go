package iterm

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/screen"
)

func (c *Client) MustReadScreen(identifier string) *screen.Screen {
	result, e := c.ReadScreen(identifier)
	errors.PanicOnError(e)

	return result
}
