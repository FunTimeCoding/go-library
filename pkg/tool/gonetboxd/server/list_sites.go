package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListSites(
	w http.ResponseWriter,
	_ *http.Request,
) {
	sites, e := s.client.Sites()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Sites(sites))
}
