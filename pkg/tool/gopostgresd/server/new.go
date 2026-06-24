package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Server {
	return &Server{
		store:    s,
		reporter: r,
	}
}
