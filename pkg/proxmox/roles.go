package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Roles() []*proxmox.Role {
	result, e := c.client.Roles(c.context)
	errors.PanicOnError(e)

	return result
}
