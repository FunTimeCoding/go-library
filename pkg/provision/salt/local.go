package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Local(
	glob string,
	function string,
	a []string,
) map[string]salt.LocalClientReturn {
	result, e := c.client.LocalClient(
		c.context,
		function,
		a,
		salt.WithGlobTarget(glob),
	)
	errors.PanicOnError(e)

	return result
}
