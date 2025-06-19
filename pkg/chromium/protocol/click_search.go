package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) ClickSearch(s string) {
	p.client.RunContext(
		p.context,
		chromedp.Click(s, chromedp.NodeVisible, chromedp.BySearch),
	)
}
