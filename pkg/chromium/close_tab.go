package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) CloseTab(identifier string) error {
	b := chromedp.FromContext(c.context).Browser

	return target.CloseTarget(
		target.ID(identifier),
	).Do(cdp.WithExecutor(c.context, b))
}
