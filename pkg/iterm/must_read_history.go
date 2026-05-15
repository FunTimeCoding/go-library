package iterm

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) MustReadHistory(
	identifier string,
	count int,
) session.History {
	result, e := c.ReadHistory(identifier, count)
	errors.PanicOnError(e)

	return result
}
