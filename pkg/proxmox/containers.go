package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Containers(n *proxmox.Node) proxmox.Containers {
	result, e := n.Containers(c.context)
	errors.PanicOnError(e)

	return result
}
