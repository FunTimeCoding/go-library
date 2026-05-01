package model_context

import (
	"context"
	"fmt"
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
			"no instance selected - use use_instance first",
		)
	}

	v, e := s.store.ListSchemas(x, instance)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf(
			"database error on %s: schemas not listed",
			instance,
		))
	}

	return response.SuccessAny(v)
}
