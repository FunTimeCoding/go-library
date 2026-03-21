package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"net/http"
	"strconv"
)

func (s *Server) entryFromQuery(r *http.Request) *store.Entry {
	id, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	errors.PanicOnError(e)
	entry, f := s.store.Get(uint(id))
	errors.PanicOnError(f)

	return entry
}
