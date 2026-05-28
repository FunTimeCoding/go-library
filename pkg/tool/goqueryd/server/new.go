package server

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"

func New(s *service.Service) *Server {
	return &Server{service: s}
}
