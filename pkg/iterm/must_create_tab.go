package iterm

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) MustCreateTab() session.Session {
	result, e := c.CreateTab()
	errors.PanicOnError(e)

	return result
}
