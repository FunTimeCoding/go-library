package route

import "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"

func New(s *store.Store) *Router {
	return &Router{store: s}
}
