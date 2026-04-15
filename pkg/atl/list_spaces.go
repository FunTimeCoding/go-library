package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListSpaces() string {
	result, e := c.client.ListSpaces(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
