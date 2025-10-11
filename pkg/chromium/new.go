package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	port int,
) *Client {
	allocator, allocatorCancel := chromedp.NewRemoteAllocator(
		context.Background(),
		locator.New(host).Port(port).Scheme(constant.Socket).String(),
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
