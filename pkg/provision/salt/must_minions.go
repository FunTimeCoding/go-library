package salt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
)

func (c *Client) MustMinions() []response.Minion {
	result, e := c.Minions()
	errors.PanicOnError(e)

	return result
}
