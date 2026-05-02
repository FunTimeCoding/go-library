package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

func New(
	s *store.Store,
	p *worker.Worker,
) *Server {
	return &Server{store: s, worker: p}
}
