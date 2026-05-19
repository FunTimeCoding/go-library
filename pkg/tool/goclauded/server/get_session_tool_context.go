package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionToolContext(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
	params server.GetSessionToolContextParams,
) {
	surround := 1

	if params.Surround != nil {
		surround = *params.Surround
	}

	results := s.claude.ToolContext(identifier, params.Filter, surround)
	var entries []server.ToolContextResult

	for _, tc := range results {
		entry := server.ToolContextResult{
			ToolName: tc.ToolName,
		}
		var before []server.SessionMessage

		for _, m := range tc.Before {
			meta := m.IsMeta
			before = append(
				before,
				server.SessionMessage{
					Role:      m.Role,
					Text:      m.Text,
					Timestamp: m.Timestamp,
					IsMeta:    &meta,
				},
			)
		}

		if len(before) > 0 {
			entry.Before = &before
		}

		var after []server.SessionMessage

		for _, m := range tc.After {
			meta := m.IsMeta
			after = append(
				after,
				server.SessionMessage{
					Role:      m.Role,
					Text:      m.Text,
					Timestamp: m.Timestamp,
					IsMeta:    &meta,
				},
			)
		}

		if len(after) > 0 {
			entry.After = &after
		}

		entries = append(entries, entry)
	}

	web.EncodeNotation(
		w,
		server.ToolContextResponse{Results: entries},
	)
}
