package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	return v.Snapshot(name).Delete(c.context)
}
