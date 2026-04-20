package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type listSchemasArguments struct{}

func (s *Server) listSchemas(
	x context.Context,
	_ mcp.CallToolRequest,
	_ listSchemasArguments,
) (*mcp.CallToolResult, error) {
	instance, ok := s.activeInstance(x)

	if !ok {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	v, e := s.store.ListSchemas(x, instance)

	if e != nil {
		return response.Fail("list schemas: %v", e)
	}

	return response.SuccessAny(v)
}
