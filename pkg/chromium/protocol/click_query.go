package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) ClickQuery(selector string) {
	p.client.RunContext(
		p.context,
		chromedp.Click(selector, chromedp.NodeVisible, chromedp.ByQuery),
	)
}
