package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
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

	result := make([]server.TunnelTermination, 0, len(terminations))

	for _, t := range terminations {
		entry := server.TunnelTermination{Identifier: t.Identifier}

		if t.Role != "" {
			entry.Role = &t.Role
		}

		if t.TerminationType != "" {
			entry.TerminationType = &t.TerminationType
		}

		if t.TerminationIdentifier != 0 {
			entry.TerminationIdentifier = &t.TerminationIdentifier
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
