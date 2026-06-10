package chromium

import (
	"context"
	"sync"
)

type Client struct {
	host            string
	port            int
	allocator       context.Context
	allocatorCancel context.CancelFunc
	context         context.Context
	cancel          context.CancelFunc
	targets         map[string]context.Context
	mutex           sync.Mutex
}
