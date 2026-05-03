package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) AllocateStat(stat string) string {
	result, e := c.client.AllocateStat(
		c.context,
		client.AllocateStatParamsStat(stat),
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
