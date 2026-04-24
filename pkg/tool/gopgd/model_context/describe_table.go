package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) describeTable(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DescribeTable,
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

	v, e := s.store.DescribeTable(x, instance, schema, a.Table)

	if e != nil {
		return response.Fail("describe table: %v", e)
	}

	return response.SuccessAny(v)
}
