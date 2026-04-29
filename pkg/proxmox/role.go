package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Role(name string) (*proxmox.Permission, error) {
	if name == "" {
		return nil, fmt.Errorf("role name is required")
	}

	result, e := c.client.Role(c.context, name)

	if e != nil {
		return nil, e
	}

	return &result, nil
}
