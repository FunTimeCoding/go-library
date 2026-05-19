package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) summarize(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Summarize)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	body, e := q.RequireString(constant.Body)

	if e != nil {
		return response.Fail("body is required: %v", e)
	}

	s.service.Store.UpsertSummary(c.SessionIdentifier, c.Name, body)
	s.service.Store.UpsertEvent(
		c.SessionIdentifier,
		constant.Summarize,
		c.Name,
		"",
		body,
	)
	s.pushSummary(c.Name, body)

	return response.Success("Session summary recorded for %s", c.Name)
}
