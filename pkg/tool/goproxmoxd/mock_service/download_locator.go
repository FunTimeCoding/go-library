package mock_service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"

func (s *Service) DownloadLocator(
	_ face.ProxmoxClient,
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
) error {
	return nil
}
