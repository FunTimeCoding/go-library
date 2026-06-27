package server

import (
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func findMachine(
	c face.ProxmoxClient,
	identifier int64,
	node *string,
) (*proxmox.VirtualMachine, error) {
	if node != nil && *node != "" {
		n, e := c.Node(*node)

		if e != nil {
			return nil, e
		}

		return c.Machine(n, int(identifier))
	}

	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	for _, ns := range nodes {
		n, e := c.Node(ns.Node)

		if e != nil {
			continue
		}

		machines, e := c.Machines(n)

		if e != nil {
			continue
		}

		for _, listed := range machines {
			if int64(listed.VMID) == identifier {
				return c.Machine(n, int(identifier))
			}
		}
	}

	return nil, not_found.New("machine", identifier)
}
