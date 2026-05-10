package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	"net/http"
)

func (s *Server) entryFromQuery(r *http.Request) *entry.Entry {
	result, f := s.store.Get(
		strings.ToUnsignedIntegerStrict(
			r.URL.Query().Get(constant.Identifier),
		),
	)
	errors.PanicOnError(f)

	return result
}
