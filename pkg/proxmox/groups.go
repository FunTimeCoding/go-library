package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Groups() proxmox.Groups {
	result, e := c.client.Groups(c.context)
	errors.PanicOnError(e)

	return result
}
