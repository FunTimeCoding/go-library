package chromium

import (
	"context"
	"sync"
)

type Client struct {
	host            string
	port            int
	context         context.Context
	cancel          context.CancelFunc
	allocatorCancel context.CancelFunc
	targets         map[string]context.Context
	mu              sync.Mutex
}
