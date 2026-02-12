package proxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/proxmox/constant"
	"github.com/funtimecoding/go-library/pkg/proxmox/node_status"
)

func (c *Client) NodeStatus(name string) *node_status.Status {
	errors.PanicOnEmpty(name, constant.Name)
	n := &node_status.Payload{}
	Get(c, fmt.Sprintf("/nodes/%s/status", name), &n.Data)

	return n.Data
}
