package model_context

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server       *server.MCPServer
	client       *sentry.Client
	organization string
	reporter     face.Reporter
}
