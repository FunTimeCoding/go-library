package chromium

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Targets() []*target.Info {
	c.reconnectIfNeeded()
	result, e := chromedp.Targets(c.context)
	errors.PanicOnError(e)

	return result
}
