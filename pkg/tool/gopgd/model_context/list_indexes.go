package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/model_context/argument"
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
		return s.captureFail(
			e,
			fmt.Sprintf(
				"database error on %s: indexes for %s.%s not listed",
				instance,
				schema,
				a.Table,
			),
		)
	}

	return response.SuccessAny(v)
}
