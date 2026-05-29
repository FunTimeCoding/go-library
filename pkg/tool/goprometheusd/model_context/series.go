package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) series(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Series,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	v, e := s.service.Series(instance)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf("series failed on %s", instance),
		)
	}

	return response.SuccessAny(v)
}
