package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/model_context/argument"
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
		return s.captureFail(
			e,
			fmt.Sprintf(
				"database error on %s: tables in %s not listed",
				instance,
				schema,
			),
		)
	}

	return response.SuccessAny(v)
}
