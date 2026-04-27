package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) query(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Query,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
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
