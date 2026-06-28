package mock_service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"

func (s *Service) RollbackMachineSnapshot(
	_ face.ProxmoxClient,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:rollback-snapshot", nil
}
