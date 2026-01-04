package chromium

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Reload(x context.Context) {
	errors.PanicOnError(chromedp.Run(x, page.Reload()))
	errors.PanicOnError(
		chromedp.Run(
			x,
			chromedp.WaitReady(constant.BodySelector),
		),
	)
}
