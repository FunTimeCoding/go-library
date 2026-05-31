package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) RollbackMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	return v.Snapshot(name).Rollback(c.context)
}
