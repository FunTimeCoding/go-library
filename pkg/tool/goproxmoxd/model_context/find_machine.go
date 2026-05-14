package model_context

import "github.com/luthermonson/go-proxmox"

func (s *Server) findMachine(
	vmid int,
	nodeName string,
) (*proxmox.VirtualMachine, error) {
	if nodeName != "" {
		node, e := s.client.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return s.client.Machine(node, vmid)
	}

	nodes, e := s.client.Nodes()

	if e != nil {
		return nil, e
	}

	for _, ns := range nodes {
		node, e := s.client.Node(ns.Node)

		if e != nil {
			continue
		}

		machines, e := s.client.Machines(node)

		if e != nil {
			continue
		}

		for _, listed := range machines {
			if int(listed.VMID) == vmid {
				return s.client.Machine(node, vmid)
			}
		}
	}

	return nil, nil
}
