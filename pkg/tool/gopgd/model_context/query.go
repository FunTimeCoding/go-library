package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type queryArguments struct {
	SQL string `json:"sql"`
}

func (s *Server) query(
	x context.Context,
	_ mcp.CallToolRequest,
	a queryArguments,
) (*mcp.CallToolResult, error) {
	instance, ok := s.activeInstance(x)

	if !ok {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	if a.SQL == "" {
		return response.Fail("sql is required")
	}

	v, e := s.store.Query(x, instance, a.SQL)

	if e != nil {
		return response.Fail("query: %v", e)
	}

	return response.SuccessAny(v)
}
