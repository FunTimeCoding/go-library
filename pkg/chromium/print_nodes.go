package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) PrintNodes(
	o context.Context,
	selector string,
	attributes []string,
) {
	var nodes []*cdp.Node
	errors.PanicOnError(
		chromedp.Run(
			o,
			chromedp.Nodes(selector, &nodes, chromedp.ByQueryAll),
		),
	)

	fmt.Printf("Selector: %s\n", selector)

	for _, n := range nodes {
		for _, a := range attributes {
			fmt.Printf("  %s: %s\n", a, n.AttributeValue(a))
		}
	}
}
