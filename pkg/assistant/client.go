package assistant

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assistant/connection"
	"github.com/getsentry/sentry-go"
)

type Client struct {
	connection *connection.Connection
	hub        *sentry.Hub
	context    context.Context
	subscriber connection.Subscriber
}
