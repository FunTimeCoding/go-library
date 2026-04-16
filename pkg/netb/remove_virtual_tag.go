package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) RemoveVirtualTag(
	name string,
	tag string,
) string {
	result, e := c.client.RemoveVirtualTag(c.context, name, tag)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
