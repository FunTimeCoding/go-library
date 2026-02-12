package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Machines(n *proxmox.Node) proxmox.VirtualMachines {
	result, e := n.VirtualMachines(c.context)
	errors.PanicOnError(e)

	return result
}
