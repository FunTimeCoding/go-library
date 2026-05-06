package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Local(
	glob string,
	function string,
	a []string,
) (map[string]salt.LocalClientReturn, error) {
	return c.client.LocalClient(
		c.context,
		function,
		a,
		salt.WithGlobTarget(glob),
	)
}
