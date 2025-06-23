package protocol

import (
	"fmt"
	"github.com/chromedp/cdproto/cdp"
)

func Print(
	n *cdp.Node,
	attribute []string,
) {
	fmt.Printf("  XPath: %s\n", n.FullXPath())

	for _, a := range attribute {
		fmt.Printf("  %s: %s\n", a, n.AttributeValue(a))
	}
}
