package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	habitica *habitica.Client
}
