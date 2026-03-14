package maintenance_log

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
