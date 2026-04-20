package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type listIndexesArguments struct {
	Table  string `json:"table"`
	Schema string `json:"schema"`
}

func (s *Server) listIndexes(
	x context.Context,
	_ mcp.CallToolRequest,
	a listIndexesArguments,
) (*mcp.CallToolResult, error) {
	instance, ok := s.activeInstance(x)

	if !ok {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	if a.Table == "" {
		return response.Fail("table is required")
	}

	schema := a.Schema

	if schema == "" {
		schema = "public"
	}

	v, e := s.store.ListIndexes(x, instance, schema, a.Table)

	if e != nil {
		return response.Fail("list indexes: %v", e)
	}

	return response.SuccessAny(v)
}
