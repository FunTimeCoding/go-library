package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionsFind(
	w http.ResponseWriter,
	r *http.Request,
	params server.GetSessionsFindParams,
) {
	matches := s.claude.SessionsByTool(params.Tool)
	var result []server.FindMatch

	for _, m := range matches {
		alias := s.service.Store.GetAlias(m.Session.Identifier)
		name := m.Session.Slug

		if alias != "" {
			name = alias
		}

		result = append(
			result,
			server.FindMatch{
				Count:     m.Count,
				SessionId: m.Session.Identifier,
				Name:      name,
			},
		)
	}

	web.EncodeNotation(
		w,
		server.FindResponse{Matches: result},
	)
}
