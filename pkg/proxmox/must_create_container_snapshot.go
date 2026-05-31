package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustCreateContainerSnapshot(
	v *proxmox.Container,
	name string,
) *proxmox.Task {
	result, e := c.CreateContainerSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}
