package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustSnapshots(v *proxmox.VirtualMachine) []*proxmox.Snapshot {
	result, e := c.Snapshots(v)
	errors.PanicOnError(e)

	return result
}
