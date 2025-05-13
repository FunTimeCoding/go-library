package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) RunContext(
	o context.Context,
	a ...chromedp.Action,
) {
	errors.PanicOnError(chromedp.Run(o, a...))
}
