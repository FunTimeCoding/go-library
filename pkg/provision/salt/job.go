package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Job(identifier string) (*salt.Job, error) {
	return c.client.LookupJID(c.context, identifier)
}
