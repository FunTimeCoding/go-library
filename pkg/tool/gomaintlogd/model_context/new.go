package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
