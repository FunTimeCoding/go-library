package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) status(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Status)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	session := s.service.Store.SessionByName(c.Name)

	if session == nil {
		return response.Fail("session not found for %s", c.Name)
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("Name: %s", session.Name))

	if session.Topic != "" {
		lines = append(lines, fmt.Sprintf("Topic: %s", session.Topic))
	} else {
		lines = append(lines, "Topic: (none)")
	}

	if session.Files != "" {
		lines = append(lines, fmt.Sprintf("Files: %s", session.Files))
	}

	return response.Success(strings.Join(lines, "\n"))
}
