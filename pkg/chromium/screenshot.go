package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (c *Client) Screenshot(x context.Context) ([]byte, error) {
	var result []byte
	e := chromedp.Run(x, chromedp.CaptureScreenshot(&result))

	if e != nil {
		return nil, e
	}

	return result, nil
}
