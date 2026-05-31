package server

import "github.com/luthermonson/go-proxmox"

func (s *Server) findContainer(
	identifier int64,
	node *string,
) *proxmox.Container {
	if node != nil && *node != "" {
		n := s.client.MustNode(*node)

		return s.client.MustContainer(n, int(identifier))
	}

	for _, ns := range s.client.MustNodes() {
		n := s.client.MustNode(ns.Node)

		for _, listed := range s.client.MustContainers(n) {
			if int64(listed.VMID) == identifier {
				return s.client.MustContainer(n, int(identifier))
			}
		}
	}

	return nil
}
