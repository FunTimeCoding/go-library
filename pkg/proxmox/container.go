package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Container(
	n *proxmox.Node,
	identifier int,
) *proxmox.Container {
	result, e := n.Container(c.context, identifier)
	errors.PanicOnError(e)

	return result
}
