package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Outer(selector string) string {
	var result string
	p.client.RunContext(p.context, chromedp.OuterHTML(selector, &result))

	return result
}
