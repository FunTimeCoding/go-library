package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/mark3labs/mcp-go/server"
)

func New(c *habitica.Client) *Server {
	result := &Server{
		server: server.NewMCPServer(
			"habitica",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		habitica: c,
	}
	result.register()

	return result
}
