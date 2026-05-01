package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server     *server.MCPServer
	jira       *jira.Client
	confluence *confluence.Client
	hub        *sentry.Hub
}
