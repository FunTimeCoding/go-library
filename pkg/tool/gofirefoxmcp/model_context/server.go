package model_context

import (
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server *server.MCPServer
	client *firefox.Client
	hub    *sentry.Hub
}
