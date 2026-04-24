package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Stats() string {
	result, e := c.client.GetStats(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
