package route

import "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"

func New(s *store.Store) *Handler {
	return &Handler{store: s}
}
