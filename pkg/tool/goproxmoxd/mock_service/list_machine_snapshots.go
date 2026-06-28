package mock_service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachineSnapshots(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) ([]*proxmox.VirtualMachineSnapshot, error) {
	return nil, nil
}
