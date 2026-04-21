package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
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

	if v := r.GetString(constant.Start, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid start format: %v", e)
		}

		start = t
	}

	if v := r.GetString(constant.End, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid end format: %v", e)
		}

		end = t
	}

	records := s.store.ByTimeRange(start, end)

	return response.Success(
		"Found %d alerts:\n%s",
		len(records),
		notation.MarshalIndent(records),
	)
}
