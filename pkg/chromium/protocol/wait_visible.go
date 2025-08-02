package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) WaitVisible(s string) {
	p.client.RunContext(p.context, chromedp.WaitVisible(s))
}
