package route

import "github.com/funtimecoding/go-library/pkg/tool/gopgd/store"

func New(s *store.Store) *Router {
	return &Router{store: s}
}
