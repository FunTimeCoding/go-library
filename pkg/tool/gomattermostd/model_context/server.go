package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/monitor"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   *mattermost.Client
	monitor  *monitor.Monitor
	reporter face.Reporter
}
