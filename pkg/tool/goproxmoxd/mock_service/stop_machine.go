package mock_service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"

func (s *Service) StopMachine(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) (string, error) {
	return "mock:stop", nil
}
