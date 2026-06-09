package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateMachineSnapshot(
	identifier int64,
	name string,
	node *string,
) string {
	body := client.CreateMachineSnapshotJSONRequestBody{Name: name}
	result, e := c.client.CreateMachineSnapshot(
		c.context,
		identifier,
		&client.CreateMachineSnapshotParams{Instance: &c.instance, Node: node},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
