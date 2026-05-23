package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessions(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetSessionsParams,
) {
	limit := 0

	if p.Limit != nil {
		limit = *p.Limit
	}

	offset := 0

	if p.Offset != nil {
		offset = *p.Offset
	}

	peek := p.Peek != nil && *p.Peek
	sessions := s.service.EnrichedSessions(limit, offset)
	var result []server.SessionDetail

	for _, e := range sessions {
		d := server.SessionDetail{
			Identifier: e.Identifier,
			Timestamp:  e.Timestamp,
			Lines:      e.Lines,
		}

		if e.Name != "" {
			d.Name = &e.Name
		}

		if e.Slug != "" {
			d.Slug = &e.Slug
		}

		if e.WorkDirectory != "" {
			d.WorkDirectory = &e.WorkDirectory
		}

		if e.Branch != "" {
			d.Branch = &e.Branch
		}

		if e.Alias != "" {
			d.Alias = &e.Alias
		}

		if e.Description != "" {
			d.Description = &e.Description
		}

		if peek {
			preview := s.service.FirstUserMessage(e.Identifier)

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
