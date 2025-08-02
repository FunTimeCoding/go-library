package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Evaluate(
	s string,
	result any,
) {
	p.client.RunContext(p.context, chromedp.Evaluate(s, &result))
}
