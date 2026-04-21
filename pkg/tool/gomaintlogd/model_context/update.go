package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) update(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	id := r.GetFloat(constant.Identifier, 0)

	if id == 0 {
		return response.Fail("id is required")
	}

	e, f := s.store.Get(uint(id))

	if f != nil {
		return response.Fail("failed to get entry: %v", f)
	}

	if v := r.GetString(constant.Action, ""); v != "" {
		e.Action = v
	}

	if v := r.GetString(constant.User, ""); v != "" {
		e.User = v
	}

	if v := r.GetString(constant.System, ""); v != "" {
		e.System = v
	}

	if v := r.GetString(constant.Service, ""); v != "" {
		e.Service = v
	}

	if v := r.GetString(constant.Description, ""); v != "" {
		e.Description = v
	}

	if stamp := r.GetString(constant.Timestamp, ""); stamp != "" {
		t, g := time.Parse(time.RFC3339, stamp)

		if g != nil {
			return response.Fail(
				"invalid timestamp format: %v",
				g,
			)
		}

		e.Timestamp = t
	}

	if g := s.store.Update(e); g != nil {
		return response.Fail("failed to update entry: %v", g)
	}

	return response.Success(
		"Entry updated successfully:\n%s",
		notation.MarshalIndent(e),
	)
}
