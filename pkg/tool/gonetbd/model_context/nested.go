package model_context

import "github.com/mark3labs/mcp-go/server"

func (s *Server) Nested() *server.MCPServer {
	return s.server
}
