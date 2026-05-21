package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) send(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Send)

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	body, e := q.RequireString(constant.Body)

	if e != nil {
		return response.Fail("body is required: %v", e)
	}

	to := q.GetString(constant.To, "")
	s.service.Send(c.Callsign, to, body)

	if to != "" {
		return response.Success(fmt.Sprintf("Sent to %s", to))
	}

	return response.Success("Broadcast sent")
}
