package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) MachineSnapshots(v *proxmox.VirtualMachine) ([]*proxmox.VirtualMachineSnapshot, error) {
	return v.Snapshots(c.context)
}
