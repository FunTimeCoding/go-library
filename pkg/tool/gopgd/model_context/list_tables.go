package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listTables(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListTables,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
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
