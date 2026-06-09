package model_context

import (
	"github.com/funtimecoding/go-library/pkg/proxmox"
	upstream "github.com/luthermonson/go-proxmox"
)

func findContainer(
	c *proxmox.Client,
	identifier int,
	nodeName string,
) (*upstream.Container, error) {
	if nodeName != "" {
		node, e := c.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return c.Container(node, identifier)
	}

	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	for _, n := range nodes {
		node, f := c.Node(n.Node)

		if f != nil {
			continue
		}

		containers, g := c.Containers(node)

		if g != nil {
			continue
		}

		for _, listed := range containers {
			if int(listed.VMID) == identifier {
				return c.Container(node, identifier)
			}
		}
	}

	return nil, nil
}
