package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) CreateTab(l string) (string, error) {
	b := chromedp.FromContext(c.context).Browser
	identifier, e := target.CreateTarget(l).Do(
		cdp.WithExecutor(c.context, b),
	)

	if e != nil {
		return "", e
	}

	return string(identifier), nil
}
