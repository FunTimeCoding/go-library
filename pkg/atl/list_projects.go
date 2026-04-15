package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListProjects() string {
	result, e := c.client.ListProjects(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
