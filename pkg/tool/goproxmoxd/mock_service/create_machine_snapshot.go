package mock_service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"

func (s *Service) CreateMachineSnapshot(
	_ face.ProxmoxClient,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:create-snapshot", nil
}
