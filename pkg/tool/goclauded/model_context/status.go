package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) status(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Status)

	if c.Callsign == "" {
		return response.Fail(
			"unknown session - announce first to bind your identity",
		)
	}

	e := s.service.SessionByCallsign(c.Callsign)

	if e == nil {
		return response.Fail("session not found for %s", c.Callsign)
	}

	lines := []string{fmt.Sprintf("Name: %s", e.Name)}

	if e.Topic != "" {
		lines = append(lines, fmt.Sprintf("Topic: %s", e.Topic))
	} else {
		lines = append(lines, "Topic: (none)")
	}

	if e.Files != "" {
		lines = append(lines, fmt.Sprintf("Files: %s", e.Files))
	}

	return response.Success(join.NewLine(lines))
}
