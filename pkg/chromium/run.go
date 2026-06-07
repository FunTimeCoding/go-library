package chromium

import "github.com/chromedp/chromedp"

func (c *Client) Run(a ...chromedp.Action) {
	c.reconnectIfNeeded()
	c.RunContext(c.context, a...)
}
