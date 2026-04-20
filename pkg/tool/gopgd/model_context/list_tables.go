package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type listTablesArguments struct {
	Schema string `json:"schema"`
}

func (s *Server) listTables(
	x context.Context,
	_ mcp.CallToolRequest,
	a listTablesArguments,
) (*mcp.CallToolResult, error) {
	instance, ok := s.activeInstance(x)

	if !ok {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	schema := a.Schema

	if schema == "" {
		schema = "public"
	}

	v, e := s.store.ListTables(x, instance, schema)

	if e != nil {
		return response.Fail("list tables: %v", e)
	}

	return response.SuccessAny(v)
}
