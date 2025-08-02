package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Outer(s string) string {
	var result string
	p.client.RunContext(p.context, chromedp.OuterHTML(s, &result))

	return result
}
