package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Keys() (*salt.ListKeysReturn, error) {
	return c.client.ListKeys(c.context)
}
