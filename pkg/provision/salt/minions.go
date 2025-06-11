package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Minions() []salt.Minion {
	result, e := c.client.ListMinions(c.context)
	errors.PanicOnError(e)

	return result
}
