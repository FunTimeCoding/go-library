package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListContainerSnapshots(
	identifier int64,
	node *string,
) string {
	result, e := c.client.ListContainerSnapshots(
		c.context,
		identifier,
		&client.ListContainerSnapshotsParams{Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
