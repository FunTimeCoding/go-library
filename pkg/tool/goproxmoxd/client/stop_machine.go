package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) StopMachine(vmid int64, node *string) string {
	result, e := c.client.StopMachine(
		c.context,
		vmid,
		&client.StopMachineParams{Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
