package chromium

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
	"strings"
)

func (c *Client) Body(identifier string) string {
	i := target.ID(identifier)
	// Not canceling to leave the tab open
	o, _ := chromedp.NewContext(
		c.context,
		chromedp.WithTargetID(i),
	)
	var result string
	errors.PanicOnError(
		chromedp.Run(
			o,
			//target.ActivateTarget(i), // Steals focus
			chromedp.OuterHTML("body", &result),
		),
	)

	return strings.TrimSpace(result)
}
