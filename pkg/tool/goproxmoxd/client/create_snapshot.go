package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateSnapshot(
	vmid int64,
	name string,
	node *string,
) string {
	body := client.CreateSnapshotJSONRequestBody{Name: name}
	result, e := c.client.CreateSnapshot(
		c.context,
		vmid,
		&client.CreateSnapshotParams{Node: node},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
