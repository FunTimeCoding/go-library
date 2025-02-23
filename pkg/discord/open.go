package discord

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Open() {
	errors.PanicOnError(c.client.Open())
}
