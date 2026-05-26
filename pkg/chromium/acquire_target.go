package chromium

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) AcquireTarget(identifier string) context.Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	if x, okay := c.targets[identifier]; okay {
		return x
	}

	x, _ := chromedp.NewContext(
		c.context,
		chromedp.WithTargetID(target.ID(identifier)),
	)
	c.targets[identifier] = x

	return x
}
