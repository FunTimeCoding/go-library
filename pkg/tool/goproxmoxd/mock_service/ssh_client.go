package mock_service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"

func (s *Service) SSHClient(_ string) (face.SnippetClient, error) {
	return s.snippetClient, nil
}
