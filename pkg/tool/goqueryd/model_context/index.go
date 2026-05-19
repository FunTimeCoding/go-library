package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) index(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	results, e := s.service.IndexCollections(
		q.GetString(constant.Collection, ""),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success(notation.MarshalIndent(results))
}
