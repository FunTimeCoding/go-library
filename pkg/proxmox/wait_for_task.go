package proxmox

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) WaitForTask(
	t *proxmox.Task,
	seconds int,
) error {
	e := t.WaitFor(c.context, seconds)

	if e != nil {
		return e
	}

	if t.IsFailed {
		return fmt.Errorf("%s", t.ExitStatus)
	}

	return nil
}
