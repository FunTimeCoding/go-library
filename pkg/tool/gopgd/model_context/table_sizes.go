package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type tableSizesArguments struct {
	Schema string `json:"schema"`
}

func (s *Server) tableSizes(
	x context.Context,
	_ mcp.CallToolRequest,
	a tableSizesArguments,
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

	v, e := s.store.TableSizes(x, instance, schema)

	if e != nil {
		return response.Fail("table sizes: %v", e)
	}

	return response.SuccessAny(v)
}
