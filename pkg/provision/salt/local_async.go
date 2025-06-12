package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) LocalAsync(
	target string,
	function string,
	a []string,
) string {
	result, e := c.client.LocalClientAsync(
		c.context,
		target,
		function,
		a,
		salt.WithGlobTarget(target),
	)
	errors.PanicOnError(e)

	return result
}
