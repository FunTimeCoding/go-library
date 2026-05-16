package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTunnels(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tunnels, e := s.client.Tunnels()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Tunnels(tunnels))
}
