package maintenance_log

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maintenance_log/entry"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) add(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	action, f := r.RequireString("action")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("action is required: %v", f),
		), nil
	}

	user, f := r.RequireString("user")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("user is required: %v", f),
		), nil
	}

	e := &entry.Entry{Action: action, User: user}

	if y := r.GetString("system", ""); y != "" {
		e.System = y
	}

	if v := r.GetString("service", ""); v != "" {
		e.Service = v
	}

	if d := r.GetString("description", ""); d != "" {
		e.Description = d
	}

	if stamp := r.GetString("timestamp", ""); stamp != "" {
		t, parseFail := time.Parse(time.RFC3339, stamp)

		if parseFail != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid timestamp format: %v", parseFail),
			), nil
		}

		e.Timestamp = t
	}

	if g := s.store.Add(e); g != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("failed to add entry: %v", g),
		), nil
	}

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Entry added successfully:\n%s",
			notation.MarshallIndent(e),
		),
	), nil
}
