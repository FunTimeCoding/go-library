package protocol

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (p *Protocol) Select(
	selector string,
	index int,
) *cdp.Node {
	var result []*cdp.Node
	errors.PanicOnError(
		chromedp.Run(
			p.context,
			chromedp.Nodes(selector, &result, chromedp.ByQueryAll),
		),
	)

	if index < 0 || index >= len(result) {
		return nil
	}

	return result[index]
}
