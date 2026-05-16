package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/history"
)

func (c *Client) MustReadHistory(
	identifier string,
	count int,
) *history.History {
	result, e := c.ReadHistory(identifier, count)
	errors.PanicOnError(e)

	return result
}
