package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListPrefixes(
	w http.ResponseWriter,
	_ *http.Request,
) {
	prefixes, e := s.client.Prefixes()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Prefixes(prefixes))
}
