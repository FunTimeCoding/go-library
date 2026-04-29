package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Domains() (proxmox.Domains, error) {
	return c.client.Domains(c.context)
}
