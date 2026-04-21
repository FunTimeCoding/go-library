package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) delete(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	id := r.GetFloat(constant.Identifier, 0)

	if id == 0 {
		return response.Fail("id is required")
	}

	if e := s.store.Delete(uint(id)); e != nil {
		return response.Fail("failed to delete entry: %v", e)
	}

	return response.Success("Entry %d deleted successfully", int(id))
}
