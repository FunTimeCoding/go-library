package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) embed(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.store.Embed(s.ollama)

	if e != nil {
		return s.captureFail(e, "embed failed")
	}

	return response.Success(
		"embedded %d documents (%d chunks)",
		result.Documents,
		result.Chunks,
	)
}
