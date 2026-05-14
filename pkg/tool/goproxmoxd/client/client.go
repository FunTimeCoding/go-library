package client

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
