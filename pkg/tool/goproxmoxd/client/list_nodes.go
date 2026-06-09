package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListNodes() string {
	result, e := c.client.ListNodes(
		c.context,
		&client.ListNodesParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
