package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) WaitVisible(selector string) {
	p.client.RunContext(p.context, chromedp.WaitVisible(selector))
}
