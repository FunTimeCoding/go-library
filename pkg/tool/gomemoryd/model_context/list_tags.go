package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listTags(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	tags, e := s.service.ListTags()

	if e != nil {
		return s.captureFail(e, "failed to list tags")
	}

	return response.Success(notation.MarshalIndent(tags))
}
