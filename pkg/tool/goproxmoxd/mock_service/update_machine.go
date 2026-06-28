package mock_service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument/update_machine"
)

func (s *Service) UpdateMachine(
	_ face.ProxmoxClient,
	a *update_machine.Machine,
) error {
	return a.Validate()
}
