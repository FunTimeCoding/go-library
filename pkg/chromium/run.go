package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Run(a ...chromedp.Action) {
	errors.PanicOnError(chromedp.Run(c.context, a...))
}
