package chromium

import "github.com/chromedp/chromedp"

func (c *Client) Run(a ...chromedp.Action) {
	c.RunContext(c.context, a...)
}
