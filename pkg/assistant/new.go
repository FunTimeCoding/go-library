package assistant

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assistant/connection"
)

func New(
	host string,
	token string,
	o ...Option,
) *Client {
	result := &Client{
		connection: connection.New(host, token),
		context:    context.Background(),
	}

	for _, opt := range o {
		opt(result)
	}

	return result
}
