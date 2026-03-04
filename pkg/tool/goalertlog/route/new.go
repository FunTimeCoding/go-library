package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
)

func New(s *store.Store, p *poller.Poller) *Handler {
	return &Handler{
		store:  s,
		poller: p,
	}
}
