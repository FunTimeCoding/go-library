package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) historyCount(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	s.resolveCaller(x, constant.HistoryCount)
	count := s.service.Store.CountEvents()

	return response.Success(fmt.Sprintf("%d events", count))
}
