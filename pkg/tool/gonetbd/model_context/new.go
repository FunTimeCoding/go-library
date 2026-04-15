package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/mark3labs/mcp-go/server"
)

func New(c *netbox.Client) *Server {
	result := &Server{
		server: server.NewMCPServer(
			"netbox",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		client: c,
	}
	result.register()

	return result
}
