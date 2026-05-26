package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (c *Client) Evaluate(
	x context.Context,
	expression string,
	result any,
) error {
	return chromedp.Run(x, chromedp.Evaluate(expression, result))
}
