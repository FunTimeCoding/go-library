package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) RollbackSnapshot(vmid int64, name string, node *string) string {
	result, e := c.client.RollbackSnapshot(
		c.context,
		vmid,
		name,
		&client.RollbackSnapshotParams{Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
