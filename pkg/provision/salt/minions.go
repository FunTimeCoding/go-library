package salt

import "github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"

func (c *Client) Minions() ([]response.Minion, error) {
	return c.basic.ListMinions()
}
