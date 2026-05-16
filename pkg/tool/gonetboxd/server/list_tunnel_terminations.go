package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTunnelTerminations(
	w http.ResponseWriter,
	_ *http.Request,
) {
	terminations, e := s.client.TunnelTerminations()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.TunnelTerminations(terminations))
}
