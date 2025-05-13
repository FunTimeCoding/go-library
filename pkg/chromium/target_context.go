package chromium

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) TargetContext(identifier string) context.Context {
	result, _ := chromedp.NewContext(
		c.context,
		chromedp.WithTargetID(target.ID(identifier)),
	)

	return result
}
