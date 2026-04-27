package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSchemas(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListSchemas,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
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
