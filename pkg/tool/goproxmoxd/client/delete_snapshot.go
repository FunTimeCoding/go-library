package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) DeleteSnapshot(
	vmid int64,
	name string,
	node *string,
) string {
	result, e := c.client.DeleteSnapshot(
		c.context,
		vmid,
		name,
		&client.DeleteSnapshotParams{Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
