package server

import "github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"

func New(s *store.Store) *Server {
	return &Server{store: s}
}
