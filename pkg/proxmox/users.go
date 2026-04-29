package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Users() ([]*proxmox.User, error) {
	return c.client.Users(c.context)
}
