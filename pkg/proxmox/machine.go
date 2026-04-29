package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Machine(
	n *proxmox.Node,
	identifier int,
) (*proxmox.VirtualMachine, error) {
	return n.VirtualMachine(c.context, identifier)
}
