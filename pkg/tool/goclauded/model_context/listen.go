package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listen(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Listen)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	s.service.Store.SetListening(c.Name, true)

	return response.Success("Listening for messages. You will be woken when a message arrives.")
}
