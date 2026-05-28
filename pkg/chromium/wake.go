package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
)

func (c *Client) Wake(identifier string) error {
	b, e := c.browser()

	if e != nil {
		return e
	}

	return target.ActivateTarget(target.ID(identifier)).Do(
		cdp.WithExecutor(c.context, b),
	)
}
