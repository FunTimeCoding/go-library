package alert_log

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
