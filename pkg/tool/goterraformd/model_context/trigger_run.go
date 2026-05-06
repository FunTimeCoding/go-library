package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) triggerRun(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	e := s.runner.Trigger()

	if e != nil {
		return response.Fail(e.Error())
	}

	return response.Success("run queued")
}
