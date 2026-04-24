package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getStats(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	return response.SuccessAny(convert.Stats(s.habitica.UserStats()))
}
