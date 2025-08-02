package protocol

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) HasNodes(s string) bool {
	var nodes []*cdp.Node

	p.client.RunContext(
		p.context,
		chromedp.Nodes(s, &nodes, chromedp.AtLeast(0)),
	)

	return len(nodes) > 0
}
