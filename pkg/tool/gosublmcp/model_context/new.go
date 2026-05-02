package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *sublime.Client,
	r face.Reporter,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
		),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
