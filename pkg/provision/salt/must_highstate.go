package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustHighstate(
	target string,
) map[string]salt.LocalClientReturn {
	result, e := c.Highstate(target)
	errors.PanicOnError(e)

	return result
}
