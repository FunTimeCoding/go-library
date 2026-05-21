package conversations

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service"

func New(s *service.Service) *Server {
	return &Server{service: s}
}
