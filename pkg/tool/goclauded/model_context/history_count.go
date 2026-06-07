package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) historyCount(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	s.resolveCaller(x, constant.HistoryCount)
	count, e := s.service.CountEvents()

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success(fmt.Sprintf("%d events", count))
}
