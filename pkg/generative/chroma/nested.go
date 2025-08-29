package chroma

import "github.com/amikos-tech/chroma-go/pkg/api/v2"

func (c *Client) Nested() v2.Client {
	return c.client
}
