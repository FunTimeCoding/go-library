package chroma

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Close() {
	errors.LogClose(c.client)
}
