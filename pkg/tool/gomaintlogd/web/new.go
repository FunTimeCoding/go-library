package web

import "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"

func NewServer(s *store.Store) *Server {
	return &Server{store: s}
}
