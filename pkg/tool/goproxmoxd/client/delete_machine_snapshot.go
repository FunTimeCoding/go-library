package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) DeleteMachineSnapshot(
	identifier int64,
	name string,
	node *string,
) string {
	result, e := c.client.DeleteMachineSnapshot(
		c.context,
		identifier,
		name,
		&client.DeleteMachineSnapshotParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
