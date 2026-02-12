package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Domains() proxmox.Domains {
	result, e := c.client.Domains(c.context)
	errors.PanicOnError(e)

	return result
}
