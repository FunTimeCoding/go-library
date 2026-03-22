package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
)

func (s *Server) recentTable() g.Node {
	f := store.NewFilter()
	f.Limit = 50
	entries, e := s.store.List(f)
	errors.PanicOnError(e)

	return entriesTable(entries)
}
