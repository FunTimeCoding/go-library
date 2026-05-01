package model_context

import (
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type Server struct {
	server *server.MCPServer
	client *gitlab.Client
	hub    *sentry.Hub
}
