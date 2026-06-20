package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) WaitForTask(t *proxmox.Task, seconds int) error {
	return t.WaitFor(c.context, seconds)
}
