package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSessions(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListSessions,
) (*mcp.CallToolResult, error) {
	v, e := s.client.Sessions()

	if e != nil {
		return s.captureFail(e, "list sessions")
	}

	return response.SuccessAny(v)
}
