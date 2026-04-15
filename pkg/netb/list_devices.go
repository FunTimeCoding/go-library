package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListDevices(query *string) string {
	result, e := c.client.ListDevices(
		c.context,
		&client.ListDevicesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
