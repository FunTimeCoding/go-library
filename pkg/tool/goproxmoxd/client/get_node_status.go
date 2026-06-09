package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetNodeStatus(name string) string {
	result, e := c.client.GetNodeStatus(
		c.context,
		name,
		&client.GetNodeStatusParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
