package chromium

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) listenTargets() {
	chromedp.ListenTarget(
		c.context,
		func(v any) {
			if e, okay := v.(*target.EventTargetDestroyed); okay {
				c.mu.Lock()
				delete(c.targets, string(e.TargetID))
				c.mu.Unlock()
			}
		},
	)
}
