package base

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Server) Store() *store.Store {
	return s.store
}
