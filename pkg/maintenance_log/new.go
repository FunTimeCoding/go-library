package maintenance_log

import (
	"github.com/funtimecoding/go-library/pkg/maintenance_log/store"
	"github.com/mark3labs/mcp-go/server"
)

func New() *Server {
	result := &Server{
		server: server.NewMCPServer(
			"maintenance-log",
			"1.0.0",
			server.WithToolCapabilities(true),
		),
		store: store.New(),
	}
	result.register()

	return result
}
