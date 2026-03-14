package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) getRecentAlerts(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v := r.GetString("start", ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid start format: %v", e),
			), nil
		}

		start = t
	}

	if v := r.GetString("end", ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid end format: %v", e),
			), nil
		}

		end = t
	}

	records := s.store.ByTimeRange(start, end)

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Found %d alerts:\n%s",
			len(records),
			notation.MarshallIndent(records),
		),
	), nil
}
