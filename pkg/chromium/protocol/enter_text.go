package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) EnterText(
	s string,
	text string,
) {
	p.client.RunContext(
		p.context,
		chromedp.Click(s),
		chromedp.SendKeys(s, text),
		chromedp.KeyEvent("\r"),
	)
}
