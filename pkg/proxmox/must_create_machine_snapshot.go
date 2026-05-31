package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustCreateMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) *proxmox.Task {
	result, e := c.CreateMachineSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}
