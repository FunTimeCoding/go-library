package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) list(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	memoryType := q.GetString(constant.Type, "")
	tag := q.GetString(constant.Tag, "")
	memories, e := s.service.Store.ListMemories(memoryType, tag, true)

	if e != nil {
		return s.captureFail(e, "failed to list memories")
	}

	return response.Success(notation.MarshalIndent(memories))
}
