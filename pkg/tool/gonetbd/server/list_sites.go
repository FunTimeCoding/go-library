package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListSites(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.Sites(s.client.MustSites()))
}
