package chromium

import "context"

type Client struct {
	host            string
	port            int
	context         context.Context
	cancel          context.CancelFunc
	allocatorCancel context.CancelFunc
}
