package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) string {
	body := client.CreateContainerSnapshotJSONRequestBody{Name: name}
	result, e := c.client.CreateContainerSnapshot(
		c.context,
		identifier,
		&client.CreateContainerSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
