package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) (*proxmox.Node, error) {
	if name == "" {
		return nil, fmt.Errorf("node name is required")
	}

	return c.client.Node(c.context, name)
}
