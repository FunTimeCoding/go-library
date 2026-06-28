package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNodes(
	c face.ProxmoxClient,
) (proxmox.NodeStatuses, error) {
	return c.Nodes()
}
