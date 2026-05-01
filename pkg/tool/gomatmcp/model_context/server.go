package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server  *server.MCPServer
	client  *mattermost.Client
	monitor *monitor.Monitor
	hub     *sentry.Hub
}
