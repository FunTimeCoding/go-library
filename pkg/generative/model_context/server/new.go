package server

import (
	"context"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *server.MCPServer,
	o ...Option,
) *Server {
	result := &Server{server: s, context: context.Background()}

	for _, f := range o {
		f(result)
	}

	return result
}
