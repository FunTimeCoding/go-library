package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustMachine(
	n *proxmox.Node,
	identifier int,
) *proxmox.VirtualMachine {
	result, e := c.Machine(n, identifier)
	errors.PanicOnError(e)

	return result
}
