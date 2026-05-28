package server

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"

func New(s *service.Service) *Server {
	return &Server{service: s}
}
