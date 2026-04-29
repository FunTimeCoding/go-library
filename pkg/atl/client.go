package atl

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
