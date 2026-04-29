package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustRoles() []*proxmox.Role {
	result, e := c.Roles()
	errors.PanicOnError(e)

	return result
}
