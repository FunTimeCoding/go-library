package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/getsentry/sentry-go"
)

type Server struct {
	client  *mattermost.Client
	monitor *monitor.Monitor
	hub     *sentry.Hub
}
