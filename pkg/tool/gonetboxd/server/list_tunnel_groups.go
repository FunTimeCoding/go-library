package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTunnelGroups(
	w http.ResponseWriter,
	_ *http.Request,
) {
	groups, e := s.client.TunnelGroups()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.TunnelGroups(groups))
}
