package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustRollbackSnapshot(v *proxmox.VirtualMachine, name string) *proxmox.Task {
	result, e := c.RollbackSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}
