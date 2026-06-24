package model_context

import (
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/proxmox"
	upstream "github.com/luthermonson/go-proxmox"
)

func findMachine(
	c *proxmox.Client,
	identifier int,
	nodeName string,
) (*upstream.VirtualMachine, error) {
	if nodeName != "" {
		node, e := c.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return c.Machine(node, identifier)
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

		machines, g := c.Machines(node)

		if g != nil {
			continue
		}

		for _, listed := range machines {
			if int(listed.VMID) == identifier {
				return c.Machine(node, identifier)
			}
		}
	}

	return nil, not_found.New("machine", identifier)
}
