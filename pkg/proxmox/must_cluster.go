package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustCluster() *proxmox.Cluster {
	result, e := c.Cluster()
	errors.PanicOnError(e)

	return result
}
