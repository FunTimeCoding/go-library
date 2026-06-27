package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	proxFace "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  proxFace.Service
	reporter face.Reporter
}
