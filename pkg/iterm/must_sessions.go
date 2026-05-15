package iterm

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) MustSessions() []session.Session {
	result, e := c.Sessions()
	errors.PanicOnError(e)

	return result
}
