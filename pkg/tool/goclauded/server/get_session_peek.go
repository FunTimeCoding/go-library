package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionPeek(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	p := s.service.Peek(identifier)
	var entries []server.PeekEntry

	for _, entry := range p.Entries {
		e := server.PeekEntry{
			UserText: entry.UserText,
		}

		if entry.AssistantContext != "" {
			e.AssistantContext = &entry.AssistantContext
		}

		entries = append(entries, e)
	}

	var toolCounts []server.ToolCount

	for _, tc := range p.ToolCounts {
		toolCounts = append(
			toolCounts,
			server.ToolCount{
			Name:  tc.Name,
			Count: tc.Count,
		})
	}

	web.EncodeNotation(
		w,
		server.PeekResponse{
			LineCount:        p.LineCount,
			UserMessageCount: p.UserMessageCount,
			Entries:          entries,
			ToolCounts:       toolCounts,
			TotalToolCalls:   p.TotalToolCalls,
		},
	)
}
