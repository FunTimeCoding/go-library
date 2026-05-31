package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustContainerSnapshots(v *proxmox.Container) []*proxmox.ContainerSnapshot {
	result, e := c.ContainerSnapshots(v)
	errors.PanicOnError(e)

	return result
}
