package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureFail(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return response.CaptureFail(s.reporter, e, message)
}
