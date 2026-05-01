package model_context

import (
	"context"
	"fmt"
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
			"no instance selected - use use_instance first",
		)
	}

	if a.SQL == "" {
		return response.Fail("sql is required")
	}

	v, e := s.store.Query(x, instance, a.SQL)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf(
			"database error on %s: query not executed",
			instance,
		))
	}

	return response.SuccessAny(v)
}
