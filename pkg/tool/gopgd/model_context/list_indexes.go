package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listIndexes(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListIndexes,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
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
