package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Highstate(
	target string,
) map[string]salt.LocalClientReturn {
	return c.Local(target, "state.highstate", nil)
}
