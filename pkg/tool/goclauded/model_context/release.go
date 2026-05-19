package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) release(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Release)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	s.service.Store.LogEvent(
		c.SessionIdentifier,
		constant.Release,
		c.Name,
		"",
		"",
	)
	s.service.Store.ReleaseByName(c.Name)

	return response.Success("Released")
}
