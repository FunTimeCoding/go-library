package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	firefox "github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   firefox.FirefoxSource
	reporter face.Reporter
}
