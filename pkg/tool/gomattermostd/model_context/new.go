package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/monitor"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			version,
		),
		client:   m,
		monitor:  o,
		reporter: r,
	}
	result.register()

	return result
}
