package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetPage(identifier string) string {
	result, e := c.client.GetPage(c.context, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
