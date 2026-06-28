package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListMachineSnapshots(
	identifier int64,
	node *string,
) string {
	result, e := c.client.ListMachineSnapshotsWithResponse(
		c.context,
		identifier,
		&client.ListMachineSnapshotsParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
