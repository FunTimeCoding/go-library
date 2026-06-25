package technitium

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) SelfSigned() *Client {
	c.client = web.InsecureClient()

	return c
}
