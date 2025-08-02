package protocol

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) Select(
	selector string,
	index int,
) *cdp.Node {
	var result []*cdp.Node
	p.client.RunContext(
		p.context,
		chromedp.Nodes(selector, &result, chromedp.ByQueryAll),
	)

	if index < 0 || index >= len(result) {
		return nil
	}

	return result[index]
}
