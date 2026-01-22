package maintenance_log

import "github.com/mark3labs/mcp-go/server"

func (s *Server) Nested() *server.MCPServer {
	return s.server
}
