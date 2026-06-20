package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) CreateMachine(
	n *proxmox.Node,
	identifier int,
	options ...proxmox.VirtualMachineOption,
) (*proxmox.Task, error) {
	return n.NewVirtualMachine(c.context, identifier, options...)
}
