package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server *server.MCPServer
	store  *store.Store
	poller *poller.Poller
	hub    *sentry.Hub
}
