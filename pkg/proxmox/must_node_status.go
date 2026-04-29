package proxmox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/node_status"
)

func (c *Client) MustNodeStatus(name string) *node_status.Status {
	result, e := c.NodeStatus(name)
	errors.PanicOnError(e)

	return result
}
