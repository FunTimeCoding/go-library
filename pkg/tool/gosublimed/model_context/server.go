package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	sublime "github.com/funtimecoding/go-library/pkg/tool/gosublimed/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   sublime.SublimeSource
	reporter face.Reporter
}
