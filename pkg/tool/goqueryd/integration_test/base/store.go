package base

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"

func (s *Server) Store() *store.Store {
	return s.store
}
