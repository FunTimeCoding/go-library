package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) CreateMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	return v.NewSnapshot(c.context, name)
}
