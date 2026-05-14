package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) StopMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return v.Stop(c.context)
}
