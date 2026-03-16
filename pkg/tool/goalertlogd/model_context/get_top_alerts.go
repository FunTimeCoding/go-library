package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) getTopAlerts(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	n := r.GetInt("n", 25)
	now := time.Now()
	start := now.Add(-7 * 24 * time.Hour)
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

	records := s.store.Top(n, start, end)

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Top %d alerts:\n%s",
			len(records),
			notation.MarshalIndent(records),
		),
	), nil
}
