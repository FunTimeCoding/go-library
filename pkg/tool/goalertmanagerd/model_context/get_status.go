package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getStatus(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.GetStatus,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	v, e := s.service.Status(instance)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf("get_status failed on %s", instance),
		)
	}

	return response.SuccessAny(v)
}
