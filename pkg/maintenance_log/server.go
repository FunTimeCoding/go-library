package maintenance_log

import (
	"github.com/funtimecoding/go-library/pkg/maintenance_log/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server *server.MCPServer
	store  *store.Store
}
