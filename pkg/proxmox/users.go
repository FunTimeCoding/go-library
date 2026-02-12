package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Users() []*proxmox.User {
	result, e := c.client.Users(c.context)
	errors.PanicOnError(e)

	return result
}
