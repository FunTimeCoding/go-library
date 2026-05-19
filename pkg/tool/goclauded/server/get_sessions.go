package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessions(
	w http.ResponseWriter,
	r *http.Request,
	params server.GetSessionsParams,
) {
	sessions := s.claude.Sessions()
	peek := params.Peek != nil && *params.Peek
	var result []server.SessionDetail

	for _, e := range sessions {
		d := server.SessionDetail{
			Identifier: e.Identifier,
			Timestamp:  e.Timestamp,
			Lines:      e.Lines,
		}

		if e.Slug != "" {
			d.Slug = &e.Slug
		}

		if e.CWD != "" {
			d.Cwd = &e.CWD
		}

		if e.Branch != "" {
			d.Branch = &e.Branch
		}

		alias := s.service.Store.GetAlias(e.Identifier)

		if alias != "" {
			d.Alias = &alias
		}

		if peek {
			preview := s.claude.FirstUserMessage(e.Identifier)

			if len(preview) > 60 {
				preview = preview[:60]
			}

			if preview != "" {
				d.Preview = &preview
			}
		}

		result = append(result, d)
	}

	web.EncodeNotation(
		w,
		server.SessionListResponse{Sessions: result},
	)
}
