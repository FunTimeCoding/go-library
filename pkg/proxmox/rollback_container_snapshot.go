package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) RollbackContainerSnapshot(
	v *proxmox.Container,
	name string,
) (*proxmox.Task, error) {
	return v.Snapshot(name).Rollback(c.context, false)
}
