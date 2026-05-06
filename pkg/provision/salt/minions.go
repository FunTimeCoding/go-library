package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Minions() ([]salt.Minion, error) {
	return c.client.ListMinions(c.context)
}
