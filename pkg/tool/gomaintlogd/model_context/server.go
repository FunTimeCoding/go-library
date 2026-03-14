package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server *server.MCPServer
	store  *store.Store
}
