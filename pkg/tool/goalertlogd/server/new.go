package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

func New(
	s *store.Store,
	p *worker.Worker,
	r face.Reporter,
) *Server {
	return &Server{
		store:    s,
		worker:   p,
		reporter: r,
	}
}
