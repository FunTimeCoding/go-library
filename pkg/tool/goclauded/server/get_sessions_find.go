package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionsFind(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetSessionsFindParams,
) {
	var result []server.FindMatch

	for _, m := range s.service.SessionsByTool(p.Tool) {
		name := m.Session.Slug
		e := s.service.GetSession(m.Session.Identifier)

		if e != nil && e.Alias != nil {
			name = *e.Alias
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
