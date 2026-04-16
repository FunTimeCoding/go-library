package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateVirtualMachine(
	name string,
	cluster string,
) string {
	result, e := c.client.CreateVirtualMachine(
		c.context,
		client.CreateVirtualMachineRequest{
			Name:    name,
			Cluster: cluster,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
