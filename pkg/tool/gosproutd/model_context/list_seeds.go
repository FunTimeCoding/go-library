package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSeeds(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	seeds := s.store.Seeds()
	var result []map[string]any

	for _, v := range seeds {
		result = append(
			result,
			map[string]any{
			"name":     v.Name,
			"path":     v.Path,
			"position": v.Position,
		})
	}

	return response.SuccessAny(result)
}
