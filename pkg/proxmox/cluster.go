package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Cluster() *proxmox.Cluster {
	result, e := c.client.Cluster(c.context)
	errors.PanicOnError(e)

	return result
}
