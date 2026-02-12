package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Machine(
	n *proxmox.Node,
	identifier int,
) *proxmox.VirtualMachine {
	result, e := n.VirtualMachine(c.context, identifier)
	errors.PanicOnError(e)

	return result
}
