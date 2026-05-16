package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
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

	result := make([]server.TunnelGroup, 0, len(groups))

	for _, g := range groups {
		result = append(
			result,
			server.TunnelGroup{Identifier: g.Identifier, Name: g.Name},
		)
	}

	web.EncodeNotation(w, result)
}
