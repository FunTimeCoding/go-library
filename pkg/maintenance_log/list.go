package maintenance_log

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maintenance_log/store"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) list(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	filter := &store.Filter{}

	if y := r.GetString("system", ""); y != "" {
		filter.System = y
	}

	if v := r.GetString("service", ""); v != "" {
		filter.Service = v
	}

	if u := r.GetString("user", ""); u != "" {
		filter.User = u
	}

	if since := r.GetString("since", ""); since != "" {
		i, parseFail := time.Parse(time.RFC3339, since)

		if parseFail != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid since format: %v", parseFail),
			), nil
		}

		filter.Since = i
	}

	if until := r.GetString("until", ""); until != "" {
		u, parseFail := time.Parse(time.RFC3339, until)

		if parseFail != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("invalid until format: %v", parseFail),
			), nil
		}

		filter.Until = u
	}

	if l := r.GetFloat("limit", 0); l > 0 {
		filter.Limit = int(l)
	}

	entries, e := s.store.List(filter)

	if e != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf("failed to list entries: %v", e),
		), nil
	}

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Found %d entries:\n%s",
			len(entries),
			notation.MarshallIndent(entries),
		),
	), nil
}
