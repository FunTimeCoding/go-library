package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Storages(n *proxmox.Node) (proxmox.Storages, error) {
	return n.Storages(c.context)
}
