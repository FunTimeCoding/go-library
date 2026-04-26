package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	"net/http"
	"strconv"
)

func (s *Server) entryFromQuery(r *http.Request) *entry.Entry {
	id, e := strconv.ParseUint(r.URL.Query().Get(constant.Identifier), 10, 64)
	errors.PanicOnError(e)
	result, f := s.store.Get(uint(id))
	errors.PanicOnError(f)

	return result
}
