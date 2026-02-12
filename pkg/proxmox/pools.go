package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Pools() []*proxmox.Pool {
	result, e := c.client.Pools(c.context)
	errors.PanicOnError(e)

	return result
}
