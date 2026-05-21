package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetResolve(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetResolveParams,
) {
	r := s.service.ResolveSessionIdentifier(p.Query)

	if r.Ambiguous() {
		var matches []server.ResolveMatch

		for _, m := range r.Matches {
			entry := server.ResolveMatch{
				Identifier: m.Identifier,
				Field:      m.Field,
			}

			if m.Name != "" {
				entry.Name = &m.Name
			}

			if m.Alias != "" {
				entry.Alias = &m.Alias
			}

			matches = append(matches, entry)
		}

		w.WriteHeader(http.StatusConflict)
		web.EncodeNotation(
			w,
			server.ResolveAmbiguousResponse{Matches: matches},
		)

		return
	}

	if !r.Found() {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		server.ResolveResponse{Identifier: r.Identifier()},
	)
}
