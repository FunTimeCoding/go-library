package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) RollbackSnapshot(v *proxmox.VirtualMachine, name string) (*proxmox.Task, error) {
	return v.SnapshotRollback(c.context, name)
}
