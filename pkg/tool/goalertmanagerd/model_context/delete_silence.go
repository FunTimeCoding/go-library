package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteSilence(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteSilence,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	if a.ID == "" {
		return response.Fail("id is required")
	}

	e := s.service.DeleteSilence(instance, a.ID)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("silence expired: %s", a.ID)
}
