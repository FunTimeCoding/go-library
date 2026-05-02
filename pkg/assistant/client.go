package assistant

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assistant/connection"
	"github.com/funtimecoding/go-library/pkg/face"
)

type Client struct {
	connection *connection.Connection
	reporter   face.Reporter
	context    context.Context
	subscriber connection.Subscriber
}
