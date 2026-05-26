package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
)

func (c *Client) Navigate(
	x context.Context,
	l string,
) error {
	return chromedp.Run(
		x,
		chromedp.Navigate(l),
		chromedp.WaitReady(constant.BodySelector),
	)
}
