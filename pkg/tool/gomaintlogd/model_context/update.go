package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) update(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	id := r.GetFloat("id", 0)

	if id == 0 {
		return mcp.NewToolResultError("id is required"), nil
	}

	e, f := s.store.Get(uint(id))

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("failed to get entry: %v", f),
		), nil
	}

	if v := r.GetString("action", ""); v != "" {
		e.Action = v
	}

	if v := r.GetString("user", ""); v != "" {
		e.User = v
	}

	if v := r.GetString("system", ""); v != "" {
		e.System = v
	}

	if v := r.GetString("service", ""); v != "" {
		e.Service = v
	}

	if v := r.GetString("description", ""); v != "" {
		e.Description = v
	}

	if stamp := r.GetString("timestamp", ""); stamp != "" {
		t, g := time.Parse(time.RFC3339, stamp)

		if g != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid timestamp format: %v", g),
			), nil
		}

		e.Timestamp = t
	}

	if g := s.store.Update(e); g != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("failed to update entry: %v", g),
		), nil
	}

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Entry updated successfully:\n%s",
			notation.MarshalIndent(e),
		),
	), nil
}
