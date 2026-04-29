package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
)

func New(
	s *store.Store,
	p *poller.Poller,
) *Server {
	return &Server{
		store:  s,
		poller: p,
	}
}
