package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) User(name string) (*proxmox.User, error) {
	if name == "" {
		return nil, fmt.Errorf("user name is required")
	}

	return c.client.User(c.context, name)
}
