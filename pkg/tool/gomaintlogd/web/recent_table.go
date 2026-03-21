package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
)

func (s *Server) recentTable() g.Node {
	entries, e := s.store.List(&store.Filter{Limit: 50})
	errors.PanicOnError(e)

	return entriesTable(entries)
}
