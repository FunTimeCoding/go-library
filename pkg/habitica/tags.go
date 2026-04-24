package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Tags() string {
	result, e := c.client.GetTags(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
