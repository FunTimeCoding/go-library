package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Nodes() proxmox.NodeStatuses {
	result, e := c.client.Nodes(c.context)
	errors.PanicOnError(e)

	return result
}
