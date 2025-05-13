package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func New(
	host string,
	port int,
) *Client {
	allocator, allocatorCancel := chromedp.NewRemoteAllocator(
		context.Background(),
		fmt.Sprintf("ws://%s:%d", host, port),
	)
	c, cancel := chromedp.NewContext(allocator)

	return &Client{
		host:            host,
		port:            port,
		context:         c,
		cancel:          cancel,
		allocatorCancel: allocatorCancel,
	}
}
