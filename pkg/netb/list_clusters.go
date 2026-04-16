package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListClusters() string {
	result, e := c.client.ListClusters(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
