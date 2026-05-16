package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
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

	result := make([]server.Tunnel, 0, len(tunnels))

	for _, t := range tunnels {
		entry := server.Tunnel{Identifier: t.Identifier, Name: t.Name}

		if t.Raw.Encapsulation.Value != nil {
			entry.Encapsulation = new(string(*t.Raw.Encapsulation.Value))
		}

		if t.Raw.Group.IsSet() && t.Raw.Group.Get() != nil {
			entry.Group = new(t.Raw.Group.Get().GetName())
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
