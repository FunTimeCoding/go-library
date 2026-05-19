package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moment(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Moment)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	line, e := q.RequireString(constant.Line)

	if e != nil {
		return response.Fail("line is required: %v", e)
	}

	s.service.Store.LogEvent(
		c.SessionIdentifier,
		constant.Moment,
		c.Name,
		"",
		line,
	)
	s.service.ForwardImpression(line, c.Name)

	return response.Success("moment captured")
}
