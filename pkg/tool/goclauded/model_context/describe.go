package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) describe(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Describe)
	description, e := q.RequireString(constant.Description)

	if e != nil {
		return response.Fail("description is required: %v", e)
	}

	target := q.GetString(constant.Target, "")
	var sessionIdentifier string

	if target != "" {
		sessionIdentifier = s.service.Store.ResolveSessionIdentifier(target)

		if sessionIdentifier == "" {
			return response.Fail("session not found: %s", target)
		}
	} else if c.Name != "" {
		sessionIdentifier = c.SessionIdentifier
	}

	if sessionIdentifier == "" {
		return response.Fail(
			"could not resolve session - announce first or specify a target",
		)
	}

	s.service.Store.SetDescription(sessionIdentifier, description)

	return response.Success(
		fmt.Sprintf("Description set for %s", target),
	)
}
