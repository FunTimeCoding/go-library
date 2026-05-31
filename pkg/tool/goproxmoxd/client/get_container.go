package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetContainer(
	identifier int64,
	node *string,
) string {
	result, e := c.client.GetContainer(
		c.context,
		identifier,
		&client.GetContainerParams{Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
