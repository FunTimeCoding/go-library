package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) LocalAsync(
	target string,
	function string,
	a []string,
) (string, error) {
	return c.client.LocalClientAsync(
		c.context,
		target,
		function,
		a,
		salt.WithGlobTarget(target),
	)
}
