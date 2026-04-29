package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Roles() ([]*proxmox.Role, error) {
	return c.client.Roles(c.context)
}
