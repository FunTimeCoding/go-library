package chromium

import (
	"fmt"
	"github.com/chromedp/chromedp"
)

func (c *Client) browser() (*chromedp.Browser, error) {
	if c.context.Err() != nil {
		if e := c.reconnect(); e != nil {
			return nil, e
		}
	}

	x := chromedp.FromContext(c.context)

	if x == nil {
		return nil, fmt.Errorf("chromedp context is nil")
	}

	if x.Browser != nil {
		return x.Browser, nil
	}

	_, e := chromedp.Targets(c.context)

	if e != nil {
		return nil, fmt.Errorf("browser reconnect: %w", e)
	}

	if x.Browser == nil {
		return nil, fmt.Errorf("browser is nil after reconnect")
	}

	return x.Browser, nil
}
