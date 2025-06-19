package protocol

import (
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (p *Protocol) PrintNode(
	selector string,
	attribute []string,
) {
	var result []*cdp.Node
	errors.PanicOnError(
		chromedp.Run(
			p.context,
			chromedp.Nodes(selector, &result, chromedp.ByQueryAll),
		),
	)
	fmt.Printf("Selector: %s\n", selector)

	for i, n := range result {
		fmt.Printf("Index: %d\n", i)
		fmt.Printf("  XPath: %s\n", n.FullXPath())

		for _, a := range attribute {
			fmt.Printf("  %s: %s\n", a, n.AttributeValue(a))
		}
	}
}
