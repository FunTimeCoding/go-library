package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) Snapshots(v *proxmox.VirtualMachine) ([]*proxmox.Snapshot, error) {
	return v.Snapshots(c.context)
}
