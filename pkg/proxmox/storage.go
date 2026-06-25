package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Storage(
	n *proxmox.Node,
	name string,
) (*proxmox.Storage, error) {
	if name == "" {
		return nil, fmt.Errorf("storage name is required")
	}

	return n.Storage(c.context, name)
}
