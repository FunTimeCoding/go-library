package netb

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
