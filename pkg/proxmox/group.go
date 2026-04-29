package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Group(name string) (*proxmox.Group, error) {
	if name == "" {
		return nil, fmt.Errorf("group name is required")
	}

	return c.client.Group(c.context, name)
}
