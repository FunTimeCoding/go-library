package habitica

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
)

type Client struct {
	context context.Context
	client  *client.Client
}
