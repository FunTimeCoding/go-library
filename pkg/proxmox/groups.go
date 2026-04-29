package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Groups() (proxmox.Groups, error) {
	return c.client.Groups(c.context)
}
