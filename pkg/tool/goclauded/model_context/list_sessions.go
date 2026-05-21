package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSessions(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	s.resolveCaller(x, constant.ListSessions)
	limit := int(q.GetFloat(constant.Limit, 20))
	offset := int(q.GetFloat(constant.Offset, 0))
	sessions := s.service.EnrichedSessions(limit, offset)

	if len(sessions) == 0 {
		return response.Success("No sessions found.")
	}

	var lines []string

	for _, e := range sessions {
		name := e.Slug

		if e.Alias != "" {
			name = e.Alias
		}

		if name == "" {
			name = e.Identifier[:8]
		}

		line := fmt.Sprintf(
			"%s  %s  %d lines",
			e.Identifier[:8],
			name,
			e.Lines,
		)

		if e.TurnCount > 0 {
			line = fmt.Sprintf("%s  %d turns", line, e.TurnCount)
		}

		if e.Active {
			line = fmt.Sprintf("%s  (active)", line)
		}

		if e.Description != "" {
			line = fmt.Sprintf("%s\n  %s", line, e.Description)
		}

		lines = append(lines, line)
	}

	return response.Success(join.NewLine(lines))
}
