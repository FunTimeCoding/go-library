package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustStartMachine(v *proxmox.VirtualMachine) *proxmox.Task {
	result, e := c.StartMachine(v)
	errors.PanicOnError(e)

	return result
}
