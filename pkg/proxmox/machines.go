package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Machines(n *proxmox.Node) (proxmox.VirtualMachines, error) {
	return n.VirtualMachines(c.context)
}
