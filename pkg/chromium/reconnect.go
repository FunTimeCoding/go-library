package chromium

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"log/slog"
)

func (c *Client) reconnect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.context.Err() == nil {
		return nil
	}

	c.cancel()
	clear(c.targets)
	fresh, cancel := chromedp.NewContext(c.allocator)

	if _, e := chromedp.Targets(fresh); e != nil {
		cancel()

		return fmt.Errorf("cdp reconnect: %w", e)
	}

	c.context = fresh
	c.cancel = cancel
	c.listenTargets()
	slog.Info("reconnected CDP context")

	return nil
}
