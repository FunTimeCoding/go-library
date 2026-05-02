package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/face"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	r face.Reporter,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
		),
		client:   m,
		monitor:  o,
		reporter: r,
	}
	result.register()

	return result
}
