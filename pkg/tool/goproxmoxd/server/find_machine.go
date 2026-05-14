package server

import "github.com/luthermonson/go-proxmox"

func (s *Server) findMachine(
	vmid int64,
	node *string,
) *proxmox.VirtualMachine {
	if node != nil && *node != "" {
		n := s.client.MustNode(*node)

		return s.client.MustMachine(n, int(vmid))
	}

	for _, ns := range s.client.MustNodes() {
		n := s.client.MustNode(ns.Node)

		for _, listed := range s.client.MustMachines(n) {
			if int64(listed.VMID) == vmid {
				return s.client.MustMachine(n, int(vmid))
			}
		}
	}

	return nil
}
