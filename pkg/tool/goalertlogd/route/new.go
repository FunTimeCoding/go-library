package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
)

func New(
	s *store.Store,
	p *poller.Poller,
) *Router {
	return &Router{
		store:  s,
		poller: p,
	}
}
