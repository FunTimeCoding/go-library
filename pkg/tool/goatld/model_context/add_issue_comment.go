package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addIssueComment(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	body, g := r.RequireString(parameter.Body)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	s.jira.Comment(key, body)

	return response.Success("comment added")
}
